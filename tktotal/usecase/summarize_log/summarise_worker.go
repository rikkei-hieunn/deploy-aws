package summarizelog

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"tktotal/configs"
	"tktotal/infrastructure/repository"
	"tktotal/model"
	"tktotal/pkg/util"
)

type summariseLogWorker struct {
	wg           *sync.WaitGroup
	logInfo      *model.OutputLog
	config       *configs.Server
	objectKeys   []string
	s3Repository repository.IS3Repository
}

//NewWorker constructor
func NewWorker(cfg *configs.Server, s3Repository repository.IS3Repository, keys []string, logInfo *model.OutputLog) IWorker {
	return &summariseLogWorker{
		wg:           &sync.WaitGroup{},
		logInfo:      logInfo,
		config:       cfg,
		objectKeys:   keys,
		s3Repository: s3Repository,
	}
}

//Start worker
func (w *summariseLogWorker) Start(ctx context.Context) []error {
	var errs []error
	for i := 0; i < len(w.objectKeys)-1; i++ {
		w.wg.Add(1)
		go func(objectKey string) {
			defer w.wg.Done()
			err := w.handleLogForEachFile(ctx, objectKey)
			if err != nil {
				errs = append(errs, err)

				return
			}
		}(w.objectKeys[i])
	}
	w.wg.Wait()
	//update new count value
	w.summariseKubunHassinInfo()
	w.summarisePortSyubetuInfo()
	err := w.exportLogToS3(ctx)
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

// handleLogForEachFile インプットログファイルごとにログを統計する
func (w *summariseLogWorker) handleLogForEachFile(ctx context.Context, objectKey string) error {
	var err error
	//var err error
	logData, err := w.s3Repository.Download(ctx, objectKey)
	if err != nil {
		return fmt.Errorf("download file s3 fail: %w ", err)
	}
	dataBytes := logData.Bytes()
	for {
		var inputLog model.InputLog
		bufferLength, line, err := bufio.ScanLines(dataBytes, true)
		if err != nil {
			return fmt.Errorf("read line fail : %w ", err)
		}
		if bufferLength == 0 {
			break
		}

		err = json.Unmarshal(line, &inputLog)
		if err != nil {
			return fmt.Errorf("unmarshal json fail : %w ", err)
		}

		err = w.processMinuteInfo(&inputLog)
		if err != nil {
			return fmt.Errorf("process log for one minute fail : %w ", err)
		}

		w.processLogPortSyubetu(&inputLog)

		w.processLogUserID(&inputLog)

		w.processLogQuoteCode(&inputLog)

		w.processLogElement(&inputLog)

		if bufferLength <= len(dataBytes) {
			dataBytes = dataBytes[bufferLength:]
		}
	}

	return nil
}

func (w *summariseLogWorker) processMinuteInfo(inputLog *model.InputLog) error {
	timestamp, err := time.Parse(time.RFC3339Nano, inputLog.Timestamp)
	if err != nil {
		return err
	}
	minute := fmt.Sprintf(model.MinuteFormat, timestamp.Hour(), timestamp.Minute())
	w.logInfo.Mu.Lock()
	w.logInfo.MinuteInfo[minute]++
	w.logInfo.Mu.Unlock()

	return nil
}

func (w *summariseLogWorker) processLogPortSyubetu(inputLog *model.InputLog) {
	if inputLog.Level == model.ErrorLevel {
		w.logInfo.Mu.Lock()
		key := inputLog.Port + model.StrokeCharacter + model.ErrorSyubetu
		w.logInfo.SyubetuInfo[key]++
		w.logInfo.Mu.Unlock()

		return
	}
	syubetu := inputLog.ReceiveHeader[model.SyubetuStartIndex:model.SyubetuEndIndex]
	var isDefinedSyubetu bool
	for i := range w.config.Suybetu {
		if syubetu == w.config.Suybetu[i] {
			w.logInfo.Mu.Lock()
			key := inputLog.Port + model.StrokeCharacter + syubetu
			w.logInfo.SyubetuInfo[key]++
			w.logInfo.Mu.Unlock()
			isDefinedSyubetu = true
		}
	}
	if !isDefinedSyubetu {
		w.logInfo.Mu.Lock()
		key := inputLog.Port + model.StrokeCharacter + model.OtherSyubetu
		w.logInfo.SyubetuInfo[key]++
		w.logInfo.Mu.Unlock()
	}
}

func (w *summariseLogWorker) processLogUserID(inputLog *model.InputLog) {
	userID := strings.Trim(inputLog.ReceiveHeader[model.UserIDStartIndex:model.UserIDEndIndex], model.SpaceCharacter)
	w.logInfo.Mu.Lock()
	w.logInfo.UserInfo[userID]++
	w.logInfo.Mu.Unlock()
}

func (w *summariseLogWorker) processLogQuoteCode(inputLog *model.InputLog) {
	quoteCode := strings.Trim(inputLog.ReceiveHeader[model.QuoteCodeStartIndex:model.QuoteCodeEndIndex], model.SpaceCharacter)
	w.logInfo.Mu.Lock()
	w.logInfo.QuoteCodeInfo[quoteCode]++
	w.logInfo.Mu.Unlock()
}

func (w *summariseLogWorker) processLogElement(inputLog *model.InputLog) {
	elements := strings.Split(inputLog.ResponseElement, model.CommaCharacter)
	for _, element := range elements {
		w.logInfo.ElementInfo[element]++
	}
}

func (w *summariseLogWorker) summariseKubunHassinInfo() {
	for qcd, qcdCount := range w.logInfo.QuoteCodeInfo {
		kubunHasshin := util.GetKubunHassinPairFromQcd(qcd)
		w.logInfo.KubunHassinInfo[kubunHasshin] += qcdCount
	}
}
func (w *summariseLogWorker) exportLogToS3(ctx context.Context) error {
	var buffer bytes.Buffer
	//日付を出力
	buffer.WriteString(w.logInfo.Date)
	buffer.WriteString(model.NextLineCharacter)
	buffer.WriteString(model.NextLineCharacter)

	//QUOTE区分・発信元のログを出力
	buffer.WriteString(model.KubunHeader + model.CommaCharacter + model.HassinHeader)
	buffer.WriteString(model.NextLineCharacter)
	for kubunHassin, count := range w.logInfo.KubunHassinInfo {
		kubun, hassin := util.SplitTexts(kubunHassin, model.StrokeCharacter)
		buffer.WriteString(kubun + model.CommaCharacter + hassin + model.CommaCharacter + strconv.Itoa(count))
		buffer.WriteString(model.NextLineCharacter)
	}
	buffer.WriteString(model.NextLineCharacter)

	//エレメントのログを出力
	buffer.WriteString(model.ElementHeader)
	buffer.WriteString(model.NextLineCharacter)
	for elementName, count := range w.logInfo.ElementInfo {
		buffer.WriteString(elementName + model.CommaCharacter + strconv.Itoa(count))
		buffer.WriteString(model.NextLineCharacter)
	}
	buffer.WriteString(model.NextLineCharacter)

	//ユーザーIDのログを出力
	buffer.WriteString(model.UserIDHeader + model.NextLineCharacter)
	for userID, count := range w.logInfo.UserInfo {
		buffer.WriteString(userID + model.CommaCharacter + strconv.Itoa(count))
		buffer.WriteString(model.NextLineCharacter)
	}
	buffer.WriteString(model.NextLineCharacter)

	//ポートと種別のログを出力
	buffer.WriteString(model.ServerNameHeader + model.CommaCharacter + model.PortHeader + model.CommaCharacter + model.SumHeader +
		strings.Join(w.config.Suybetu, model.CommaCharacter))
	buffer.WriteString(model.NextLineCharacter)
	for i := range w.config.Port {
		buffer.WriteString(model.CommaCharacter + w.config.Port[i])
		for j := range w.config.Suybetu {
			buffer.WriteString(model.CommaCharacter + strconv.Itoa(w.logInfo.SyubetuInfo[w.config.Port[i]+model.StrokeCharacter+w.config.Suybetu[j]]))
		}
		buffer.WriteString(model.NextLineCharacter)
	}
	buffer.WriteString(model.NextLineCharacter)

	//分間のログを出力
	buffer.WriteString(model.MinuteHeader + model.NextLineCharacter)
	for i := range model.Minutes {
		hour, minute := util.SplitTexts(model.Minutes[i], model.ColonCharacter)
		buffer.WriteString(fmt.Sprintf(model.MinuteFormatFull, hour, minute) + model.CommaCharacter +
			strconv.Itoa(w.logInfo.MinuteInfo[hour+model.ColonCharacter+minute]))
		buffer.WriteString(model.NextLineCharacter)
	}

	//ファイルパスを取得
	path := w.config.OutputLogPath + w.logInfo.WeekDay + model.CSVExtension
	err := w.s3Repository.Upload(ctx, path, buffer)
	if err != nil {
		return fmt.Errorf("upload file s3 fail : %w", err)
	}

	return nil
}

func (w *summariseLogWorker) summarisePortSyubetuInfo() {
	for portSyubetu, count := range w.logInfo.SyubetuInfo {
		port, _ := util.SplitTexts(portSyubetu, model.StrokeCharacter)
		key := port + model.StrokeCharacter + model.SumSyubetu
		w.logInfo.SyubetuInfo[key] += count
	}
}
