package summarizelog

import (
	"context"
	"github.com/rs/zerolog/log"
	"sync"
	"tktotal/configs"
	"tktotal/infrastructure/repository"
	"tktotal/model"
	"tktotal/pkg/util"
)

type service struct {
	wg           *sync.WaitGroup
	config       *configs.Server
	s3Repository repository.IS3Repository
}

// NewService construct
func NewService(cfg *configs.Server, s3Repository repository.IS3Repository) ILogSummarizer {
	return &service{
		wg:           &sync.WaitGroup{},
		config:       cfg,
		s3Repository: s3Repository,
	}
}

func (s *service) Start(ctx context.Context) error {
	//Loop through date range
	for _, date := range s.config.Dates {
		var keys []string
		dateStr := date.Format(model.DateFormatWithoutStroke)

		inforLogS3Path := s.config.InfoLogPath + dateStr + model.StrokeCharacter
		errorLogS3Path := s.config.ErrorLogPath + dateStr + model.StrokeCharacter

		inforObjKeys, err := s.s3Repository.GetObjectKeys(ctx, inforLogS3Path)
		if err != nil {
			return err
		}
		keys = append(keys, inforObjKeys...)

		errObjKeys, err := s.s3Repository.GetObjectKeys(ctx, errorLogS3Path)
		if err != nil {
			return err
		}
		keys = append(keys, errObjKeys...)

		outputLog := &model.OutputLog{
			Date:            dateStr,
			WeekDay:         util.GetDayName(date.Weekday()),
			ElementInfo:     make(map[string]int),
			UserInfo:        make(map[string]int),
			QuoteCodeInfo:   make(map[string]int),
			MinuteInfo:      make(map[string]int),
			SyubetuInfo:     make(map[string]int),
			KubunHassinInfo: make(map[string]int),
		}
		if len(keys) == 0 {
			log.Info().Msgf("list keys is empty in %s ", dateStr)

			continue
		}

		worker := NewWorker(s.config, s.s3Repository, keys, outputLog)
		s.wg.Add(1)
		go func(worker IWorker) {
			defer s.wg.Done()
			errs := worker.Start(ctx)
			//TODO confirm format log senju
			if len(errs) != 0 {
				log.Error().Msgf("worker working fail with date : %s", dateStr)
			}
		}(worker)
	}

	s.wg.Wait()

	return nil
}
