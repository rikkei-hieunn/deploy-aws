package check_messages

import (
	"context"
	"fmt"
	"math"
	"message-receive-check/configs"
	"message-receive-check/infrastructure/repository"
	"message-receive-check/model"
	"os"
	"strconv"
)

type checkMessagesWorker struct {
	config       *configs.Server
	s3Repository repository.IS3Repository
}

// NewWorker constructor create a new worker
func NewWorker(config *configs.Server, s3Repo repository.IS3Repository) IWorker {
	return &checkMessagesWorker{
		config:       config,
		s3Repository: s3Repo,
	}
}

// Start worker
func (w *checkMessagesWorker) Start(ctx context.Context, fileName string) error {
	// download file the first kei
	var err error
	var kei1CountBytes, kei2CountBytes []byte

	if w.config.DevelopEnvironment {
		kei1CountBytes, err = os.ReadFile(w.config.MessageCountLogKei1Object + fileName + model.LogFileExtension)
	} else {
		kei1CountBytes, err = w.s3Repository.Download(w.config.MessageCountLogKei1Object + fileName + model.LogFileExtension)
	}
	if err != nil {
		return fmt.Errorf("cannot download file %s", fileName)
	}

	// download file the second kei
	if w.config.DevelopEnvironment {
		kei2CountBytes, err = os.ReadFile(w.config.MessageCountLogKei2Object + fileName + model.LogFileExtension)
	} else {
		kei2CountBytes, err = w.s3Repository.Download(w.config.MessageCountLogKei2Object + fileName + model.LogFileExtension)
	}
	if err != nil {
		return fmt.Errorf("cannot download file %s", fileName)
	}

	// parse count from string to float
	kei1Count, err := strconv.ParseFloat(string(kei1CountBytes), 32)
	if err != nil {
		return fmt.Errorf("cannot parse to string %v", kei1Count)
	}
	kei2Count, err := strconv.ParseFloat(string(kei2CountBytes), 32)
	if err != nil {
		return fmt.Errorf("cannot parse to string %v", kei2Count)
	}

	smallerNum := kei1Count
	if kei2Count < kei1Count {
		smallerNum = kei1Count
	}

	difference := math.Abs(kei2Count-kei1Count) / smallerNum * 100
	if difference > float64(w.config.NumberPercentAlert) {
		return fmt.Errorf("logic group: %s over %d", fileName, w.config.NumberPercentAlert)
	}

	return nil
}
