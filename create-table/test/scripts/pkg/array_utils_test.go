package pkg

import (
	"create-table/configs"
	"create-table/model"
	"create-table/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var targets = []configs.TargetCreateTable{
	{
		QKbn:     "E",
		Sndc:     "T",
		DataType: model.HigaDataString,
	},
	{
		QKbn:     "E",
		Sndc:     "T",
		DataType: model.TickDataString,
	},
	{
		QKbn:     "E",
		Sndc:     "T",
		DataType: model.KehaiDataString,
	},
	{
		QKbn:     "E",
		Sndc:     "T",
		DataType: model.JishouDataString,
	},
	{
		QKbn:     "E",
		Sndc:     "T",
		DataType: model.OneMinuteData,
	},
}

func TestContain(t *testing.T) {
	tests := []struct {
		name     string
		kubun    string
		hassin   string
		dataType string
		expect   bool
	}{
		{
			name:     "Kubun missing",
			kubun:    "",
			hassin:   "T",
			dataType: model.TickDataString,
			expect:   false,
		},
		{
			name:     "Hassin missing",
			kubun:    "E",
			hassin:   "",
			dataType: model.TickDataString,
			expect:   false,
		},
		{
			name:     "Data type missing",
			kubun:    "E",
			hassin:   "T",
			dataType: "",
			expect:   false,
		},
		{
			name:     "All arguments missing",
			kubun:    "",
			hassin:   "",
			dataType: "",
			expect:   false,
		},
		{
			name:     "Kubun not matched",
			kubun:    "@",
			hassin:   "T",
			dataType: model.HigaDataString,
			expect:   false,
		},
		{
			name:     "Hassin not matched",
			kubun:    "E",
			hassin:   "O",
			dataType: model.TickDataString,
			expect:   false,
		},
		{
			name:     "Data type not matched",
			kubun:    "E",
			hassin:   "T",
			dataType: "9",
			expect:   false,
		},
		{
			name:     "All arguments not matched",
			kubun:    "@",
			hassin:   "O",
			dataType: "9",
			expect:   false,
		},
		{
			name:     "Success",
			kubun:    "E",
			hassin:   "T",
			dataType: model.OneMinuteData,
			expect:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.Contain(targets, tt.kubun, tt.hassin, tt.dataType)
			assert.Equal(t, result, tt.expect)
		})
	}
}
