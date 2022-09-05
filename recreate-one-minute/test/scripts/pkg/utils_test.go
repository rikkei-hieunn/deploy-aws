package pkg

import (
	"github.com/stretchr/testify/assert"
	"recreate-one-minute/model"
	"recreate-one-minute/pkg"
	"testing"
)

func Test_IsRecordExisted(t *testing.T) {
	type args struct {
		record  model.Record
		records []model.Record
	}
	recordList := []model.Record{
		{Type: "type1", Kubun: "E", Hasshin: "T", CreateDay: "2022-08-31", CreateTime: "09:00", StartIndex: "1", PathFolder: "path/folder1"},
		{Type: "type2", Kubun: "E", Hasshin: "CXJ", CreateDay: "2022-08-31", CreateTime: "09:00", StartIndex: "1", PathFolder: "path/folder1"},
	}
	tests := []struct {
		name   string
		args   args
		expect bool
	}{
		{
			name: "record founds in list record",
			args: struct {
				record  model.Record
				records []model.Record
			}{record: model.Record{
				Type:       "type2",
				Kubun:      "E",
				Hasshin:    "CXJ",
				CreateDay:  "2022-08-31",
				CreateTime: "09:00",
				StartIndex: "1",
				PathFolder: "path/folder1",
			}, records: recordList},
			expect: true,
		}, {
			name: "record is not founds in list record",
			args: struct {
				record  model.Record
				records []model.Record
			}{record: model.Record{
				Type:       "type3",
				Kubun:      "E",
				Hasshin:    "CXJ",
				CreateDay:  "2022-08-31",
				CreateTime: "09:00",
				StartIndex: "1",
				PathFolder: "path/folder1",
			}, records: recordList},
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pkg.IsRecordExisted(test.args.record, test.args.records)
			assert.Equal(t, test.expect, result)
		})
	}
}

func Test_IsItemExisted(t *testing.T) {
	type args struct {
		itemFind string
		items    []string
	}
	itemsList := []string{"@/LN", "@/TL", "E/CXJ"}
	tests := []struct {
		name   string
		args   args
		expect bool
	}{
		{
			name: "item founds in list record",
			args: struct {
				itemFind string
				items    []string
			}{itemFind: "@/TL", items: itemsList},
			expect: true,
		}, {
			name: "item is not found in list record",
			args: struct {
				itemFind string
				items    []string
			}{itemFind: "P/ERX", items: itemsList},
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pkg.IsItemExisted(test.args.itemFind, test.args.items)
			assert.Equal(t, test.expect, result)
		})
	}
}

func Test_GetTableName(t *testing.T) {
	type args struct {
		tableName    string
		kubunHasshin string
	}
	candleManagementPrefix := "candle_management"
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "can not split kubunHasshin by stroke character",
			args: struct {
				tableName    string
				kubunHasshin string
			}{tableName: candleManagementPrefix, kubunHasshin: "ET"},
			expect: model.EmptyString,
		}, {
			name: "get table name successfully",
			args: struct {
				tableName    string
				kubunHasshin string
			}{tableName: candleManagementPrefix, kubunHasshin: "E/T"},
			expect: candleManagementPrefix + model.UnderscoreCharacter + "E" + model.UnderscoreCharacter + "T",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			table := pkg.GetTableName(test.args.tableName, test.args.kubunHasshin)
			assert.Equal(t, test.expect, table)
		})
	}
}
