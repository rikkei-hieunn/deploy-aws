/*
Package loaddata implements logics load database status.
*/
package loaddata

import (
	"encoding/json"
	"fmt"
	"os"
	"show-status/configs"
	"show-status/infrastructure/repository"
	"show-status/model"
	"strings"
)

type service struct {
	config       *configs.Server
	s3Repository repository.IS3Repository
}

// NewService service constructor
func NewService(cfg *configs.Server, s3Repo repository.IS3Repository) ILoader {
	return &service{
		config:       cfg,
		s3Repository: s3Repo,
	}
}

// LoadDatabaseStatus get database status from S3 and parse to object
func (s *service) LoadDatabaseStatus() error {
	var databaseStatusData []byte
	var err error

	// Load database status
	if s.config.DevelopEnvironment {
		databaseStatusData, err = os.ReadFile(s.config.TickSystem.DatabaseStatusDefinitionObject)
	} else {
		databaseStatusData, err = s.s3Repository.Download(s.config.TickSystem.DatabaseStatusDefinitionObject)
	}
	if err != nil {
		return err
	}

	var databaseStatuses configs.GroupDatabaseStatusDefinition
	err = json.Unmarshal(databaseStatusData, &databaseStatuses)
	if err != nil {
		return err
	}

	theFirstKeiTickData := make(map[string]bool)
	theSecondtKeiTickData := make(map[string]bool)
	theFirstKeiKehaiData := make(map[string]bool)
	theSecondKeiKehaiData := make(map[string]bool)

	for _, dbStatus := range databaseStatuses.Tick {
		if !dbStatus.TheFirstKeiStatus {
			theFirstKeiTickData[dbStatus.QKbn+model.StrokeCharacter+dbStatus.Sndc] = false
		}
		if !dbStatus.TheSecondKeiStatus {
			theSecondtKeiTickData[dbStatus.QKbn+model.StrokeCharacter+dbStatus.Sndc] = false
		}
	}

	for _, dbStatus := range databaseStatuses.Kehai {
		if !dbStatus.TheFirstKeiStatus {
			theFirstKeiKehaiData[dbStatus.QKbn+model.StrokeCharacter+dbStatus.Sndc] = false
		}
		if !dbStatus.TheSecondKeiStatus {
			theSecondKeiKehaiData[dbStatus.QKbn+model.StrokeCharacter+dbStatus.Sndc] = false
		}
	}

	firstKeiTick := getMapKeys(theFirstKeiTickData)
	secondKeiTick := getMapKeys(theSecondtKeiTickData)
	firstKeiKehai := getMapKeys(theFirstKeiKehaiData)
	secondKeiKehai := getMapKeys(theSecondKeiKehaiData)

	isSuccess := true
	sbValues := strings.Builder{}
	if len(firstKeiTick) == 0 {
		sbValues.WriteString("[TICK]1系状態[通常]\n")
	} else {
		isSuccess = false
		sbValues.WriteString(fmt.Sprintf("[TICK]1系状態[異常] %d件（%s）\n", len(firstKeiTick), strings.Join(firstKeiTick, model.CommaCharacter)))
	}

	if len(secondKeiTick) == 0 {
		sbValues.WriteString("[TICK]2系状態[通常]\n")
	} else {
		isSuccess = false
		sbValues.WriteString(fmt.Sprintf("[TICK]2系状態[異常] %d件（%s）\n", len(secondKeiTick), strings.Join(secondKeiTick, model.CommaCharacter)))
	}

	if len(firstKeiKehai) == 0 {
		sbValues.WriteString("[KEHAI]1系状態[通常]\n")
	} else {
		isSuccess = false
		sbValues.WriteString(fmt.Sprintf("[KEHAI]1系状態[異常] %d件（%s）\n", len(firstKeiKehai), strings.Join(firstKeiKehai, model.CommaCharacter)))
	}

	if len(secondKeiKehai) == 0 {
		sbValues.WriteString("[KEHAI]2系状態[通常]")
	} else {
		isSuccess = false
		sbValues.WriteString(fmt.Sprintf("[KEHAI]2系状態[異常] %d件（%s）\n", len(secondKeiKehai), strings.Join(secondKeiKehai, model.CommaCharacter)))
	}

	if isSuccess {
		fmt.Println("・すべて正常の場合：")
	} else {
		fmt.Println("・一部異常の場合：")
	}
	fmt.Println(sbValues.String())

	return nil
}

// getMapKeys get list key of map
func getMapKeys(input map[string]bool) []string {
	var result []string

	for key := range input {
		result = append(result, key)
	}

	return result
}
