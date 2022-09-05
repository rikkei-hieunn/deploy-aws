package handle_request

import (
	"github.com/stretchr/testify/assert"
	"insert-calendar-infos/model"
	handlerequest "insert-calendar-infos/usecase/handle_request"
	"testing"
)

func Test_GenQueryInsertCalendar(t *testing.T) {
	tests := []struct {
		name      string
		tableName string
		expect    string
	}{
		{
			name:      "table calendar info is empty",
			tableName: model.EmptyString,
			expect:    model.EmptyString,
		},
		{
			name:      "generate insert query success",
			tableName: "calendar_infos",
			expect:    "INSERT INTO calendar_infos (CALKBN, JAPANESE_DATE, DAY_WEEK, RTYPE_CODE, SPECIAL_CODE, CLOSED_REASON) VALUES (?,?,?,?,?,?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlerequest.GenQueryInsertCalendar(tt.tableName)
			if result != nil {
				assert.Equal(t, tt.expect, result.String())
			}
		})
	}
}

func Test_GenQueryTruncateCalendar(t *testing.T) {
	tests := []struct {
		name      string
		tableName string
		expect    string
	}{
		{
			name:      "table calendar info is empty",
			tableName: model.EmptyString,
			expect:    model.EmptyString,
		},
		{
			name:      "generate truncate query success",
			tableName: "calendar_infos",
			expect:    "TRUNCATE TABLE calendar_infos",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlerequest.GenQueryTruncateCalendar(tt.tableName)
			if result != nil {
				assert.Equal(t, tt.expect, result.String())
			}
		})
	}
}
