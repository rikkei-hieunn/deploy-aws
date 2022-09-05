package insertdata

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"process-get-data/configs"
	"process-get-data/infrastructure/repository"
	"process-get-data/model"
	"time"
)

type service struct {
	config           *configs.TickSystem
	tickDBRepository repository.ITickDBRepository
}

// NewService constructor create a new service
func NewService(cfg *configs.TickSystem, tickDBRepo repository.ITickDBRepository) ITableCreator {
	return &service{
		config:           cfg,
		tickDBRepository: tickDBRepo,
	}
}

// InsertData start insert data and create cron tab
func (s *service) InsertData(ctx context.Context, quoteCodes map[string]configs.QuoteCodes, kei string) []error {
	var crontabs []*configs.CronInfo
	var errors []error

	currentDate := time.Now()
	for i, oneMinuteConfig := range model.OneMinuteConfigs {
		processDate := currentDate.AddDate(0, 0, oneMinuteConfig.CreateDay)
		systemDayOfWeek := processDate.Weekday()
		tableName := oneMinuteConfig.TableName + model.UnderscoreCharacter + processDate.Format(model.DateFormatWithoutStroke)

		if !isWorkingDay(systemDayOfWeek, &model.OneMinuteConfigs[i]) {
			continue
		}

		args := []interface{}{
			oneMinuteConfig.QKbn,
			oneMinuteConfig.Sndc,
			oneMinuteConfig.OperatorType,
			oneMinuteConfig.StartIndex,
			oneMinuteConfig.CreateTime,
			processDate.Format(model.DateFormatWithDash),
			model.DefaultFolderPath,
			oneMinuteConfig.StartTime,
			oneMinuteConfig.EndTime,
			tableName,
			model.CreatingStatusNotCreated,
			oneMinuteConfig.EndIndex,
			oneMinuteConfig.QuoteCode,
			model.NullString,
			false,
			currentDate.Format(model.DateFormatWithDash),
			currentDate.Format(model.DateFormatWithDash),
		}

		// init connection
		connectionKey := oneMinuteConfig.QKbn + model.StrokeCharacter + oneMinuteConfig.Sndc

		connectionInfo, isExists := quoteCodes[connectionKey]
		if !isExists {
			errors = append(errors, fmt.Errorf("cannot found connection key, kubun: %s, hassin: %s, start index: %s, end index: %s",
				oneMinuteConfig.QKbn, oneMinuteConfig.Sndc, oneMinuteConfig.OriginStartIndex, oneMinuteConfig.EndIndex))

			continue
		}

		err := s.tickDBRepository.InitConnection(ctx, connectionInfo.Endpoint, connectionInfo.DBName,
			oneMinuteConfig.QKbn, oneMinuteConfig.Sndc)
		if err != nil {
			errors = append(errors, fmt.Errorf("init connection fail, endpoint: %s, database name: %s", connectionInfo.Endpoint, connectionInfo.DBName))

			continue
		}

		// insert candle management data
		var candleTableName string
		kubunInsteadOf, isExists := model.KubunsInsteadOf[oneMinuteConfig.QKbn]
		if !isExists {
			candleTableName = model.CandleManagementPrefix + model.UnderscoreCharacter + oneMinuteConfig.QKbn + model.UnderscoreCharacter + oneMinuteConfig.Sndc
		} else {
			candleTableName = model.CandleManagementPrefix + model.UnderscoreCharacter + kubunInsteadOf + model.UnderscoreCharacter + oneMinuteConfig.Sndc
		}

		err = s.tickDBRepository.InsertData(ctx, fmt.Sprintf(model.QueryStringInsertCandleManagement, candleTableName), oneMinuteConfig.QKbn,
			oneMinuteConfig.Sndc, args)
		if err != nil {
			errors = append(errors, fmt.Errorf("create table fail, endpoint: %s, database name: %s", connectionInfo.Endpoint, connectionInfo.DBName))

			continue
		}

		cronInfo := s.createCronInfo(&model.OneMinuteConfigs[i], &currentDate, processDate.Format(model.DateFormatWithDash))
		crontabs = append(crontabs, cronInfo)
	}
	err := s.createCronJobs(crontabs, kei)
	if err != nil {
		errors = append(errors, fmt.Errorf("create cron tab fail, error: %w", err))
	}

	return errors
}

// isWorkingDay check data is business day
func isWorkingDay(systemDayOfWeek time.Weekday, config *configs.OneMinuteConfig) bool {
	var workDay string
	switch systemDayOfWeek {
	case time.Sunday:
		workDay = config.Sun
	case time.Monday:
		workDay = config.Mon
	case time.Tuesday:
		workDay = config.Tue
	case time.Wednesday:
		workDay = config.Wed
	case time.Thursday:
		workDay = config.Thu
	case time.Friday:
		workDay = config.Fri
	case time.Saturday:
		workDay = config.Sat
	}

	return workDay == model.WorkingFlag
}

// createCronInfo create object information for create cron tab
func (s *service) createCronInfo(config *configs.OneMinuteConfig, systemTime *time.Time, day string) *configs.CronInfo {
	hour := config.CreateTime[:2]
	minute := config.CreateTime[3:]
	dayOfMonth := systemTime.Day()
	month := int(systemTime.Month())
	dayOfWeek := int(systemTime.Weekday())

	return &configs.CronInfo{
		Hour:        hour,
		Minute:      minute,
		DayOfMonth:  dayOfMonth,
		Month:       month,
		DayOfWeek:   dayOfWeek,
		Command:     s.config.PathStartBP03,
		RequestType: model.RequestType,
		StartIndex:  config.StartIndex,
		QKBN:        config.QKbn,
		SNDC:        config.Sndc,
		CreateDate:  day,
		CreateTime:  config.CreateTime,
	}
}

// createCronJobs create cron tab
func (s *service) createCronJobs(crons []*configs.CronInfo, kei string) error {
	var buff []byte
	buff = append(buff, []byte("crontab -l > crontab_new"+model.EnterLine)...)
	for _, cron := range crons {
		buff = append(buff, []byte(fmt.Sprintf("echo \"%s %s %d %d %d %s %d %s %s %s %s %s %s\" >> crontab_new"+model.EnterLine,
			cron.Minute, cron.Hour, cron.DayOfMonth, cron.Month, cron.DayOfWeek, cron.Command, cron.RequestType,
			cron.QKBN, cron.SNDC, cron.CreateDate, cron.CreateTime, kei, cron.StartIndex))...)
	}
	buff = append(buff, []byte("crontab crontab_new")...)
	newFile, err := os.Create(s.config.PathNewCronTabs)
	if err != nil {
		return err
	}
	// write a chunk
	_, err = newFile.Write(buff)
	if err != nil {
		return err
	}
	err = newFile.Close()
	if err != nil {
		return err
	}
	_, err = exec.Command("/bin/sh", "./"+s.config.PathNewCronTabs).Output() //nolint:gosec

	return err
}
