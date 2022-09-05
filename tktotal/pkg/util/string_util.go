/*
Package util define common method
*/
package util

import (
	"strings"
	"time"
	"tktotal/model"
)

//GetKubunHassinPairFromQcd get kubun and hasshin delimiter by stroke
func GetKubunHassinPairFromQcd(quoteCode string) string {
	kubun := GetKubun(quoteCode)
	hassin := GetHassin(quoteCode)

	return kubun + model.StrokeCharacter + hassin
}

// GetHassin get hasshin from quote code, performance faster than index 1ns
func GetHassin(quoteCode string) string {
	var i int
	for i = len(quoteCode) - 1; i >= 0; i-- {
		if quoteCode[i] == 47 {
			break
		}
	}
	if i <= 0 {
		return model.EmptyString
	}

	return quoteCode[i+1:]
}

// GetKubun get kubun from full code
func GetKubun(quoteCode string) string {
	if quoteCode == model.EmptyString {
		return model.EmptyString
	}
	if quoteCode[0] == 64 && quoteCode[1] == 64 {
		return quoteCode[:2]
	}

	return quoteCode[:1]
}

// SplitTexts get two path of text after split
func SplitTexts(fulltext string, delimiter string) (string, string) {
	if fulltext == model.EmptyString && delimiter == model.EmptyString {
		return model.EmptyString, model.EmptyString
	}
	delimiterIndex := strings.Index(fulltext, delimiter)
	if delimiterIndex == -1 {
		return model.EmptyString, model.EmptyString
	}

	return fulltext[:delimiterIndex], fulltext[delimiterIndex+1:]
}

//GetDayName get day name in week
func GetDayName(weekday time.Weekday) string {
	var dayName string
	switch weekday {
	case time.Sunday:
		dayName = "SUN"
	case time.Monday:
		dayName = "MON"
	case time.Tuesday:
		dayName = "TUE"
	case time.Wednesday:
		dayName = "WED"
	case time.Thursday:
		dayName = "THU"
	case time.Friday:
		dayName = "FRI"
	case time.Saturday:
		dayName = "SAT"
	}

	return dayName
}
