// Package pkg for store helper func
package pkg

import (
	"recreate-one-minute/model"
	"reflect"
	"strings"
)

// IsRecordExisted func check record existed
func IsRecordExisted(record model.Record, records []model.Record) bool {
	for i := range records {
		if reflect.DeepEqual(record, records[i]) {
			return true
		}
	}

	return false
}

// IsItemExisted func check record existed
func IsItemExisted(item string, items []string) bool {
	if len(items) == 0 {
		return false
	}
	for _, value := range items {
		if value == item {
			return true
		}
	}
	return false
}

//GetTableName get table name
func GetTableName(tableName, kubunHasshin string) string {
	temp := strings.Split(kubunHasshin, model.StrokeCharacter)
	if len(temp) == 0 {
		return model.EmptyString
	}
	kubun := temp[0]
	hasshin := temp[0]

	kubunHasshinInsteadOf := model.KubunInsteadOf[kubun]
	if kubunHasshinInsteadOf != model.EmptyString {
		kubun = kubunHasshinInsteadOf
	}

	return tableName + model.UnderscoreCharacter + kubun + model.UnderscoreCharacter + hasshin
}
