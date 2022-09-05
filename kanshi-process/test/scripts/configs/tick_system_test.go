package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"kanshi-process/configs"
	"kanshi-process/model"
	"testing"
)

func Test_TickSystem_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect error
	}{
		{
			name: "TK_SYSTEM_REQUEST_KINOU is empty",
			args: configs.TickSystem{
				RequestKinou:     model.EmptyString,
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_KINOU required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_KANRI_ID is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   model.EmptyString,
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_KANRI_ID required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_USER_ID is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    model.EmptyString,
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_USER_ID required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_SYUBETU is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   model.EmptyString,
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_SYUBETU required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_QUOTE_CODE is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: model.EmptyString,
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_QUOTE_CODE required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_FROM_DATE is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  model.EmptyString,
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_FROM_DATE required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_TO_DATE is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    model.EmptyString,
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_TO_DATE required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_FUNASI is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    model.EmptyString,
				RequestKikan:     "1",
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_FUNASI required"),
		},
		{
			name: "TK_SYSTEM_REQUEST_KIKAN is empty",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     model.EmptyString,
			},
			expect: errors.New("system TK_SYSTEM_REQUEST_KIKAN required"),
		},
		{
			name: "validate success",
			args: configs.TickSystem{
				RequestKinou:     "01",
				RequestKanriID:   "KNR_ID",
				RequestUserID:    "Tick_Test",
				RequestSyubetu:   "2101",
				RequestQuoteCode: "XJPY/4",
				RequestFromDate:  "00000000",
				RequestToDate:    "00000000",
				RequestFunasi:    "1",
				RequestKikan:     "1",
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			assert.Equal(t, result, tt.expect)
		})
	}
}
