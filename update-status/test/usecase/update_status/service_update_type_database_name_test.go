package update_status

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"update-status/configs"
	"update-status/model"
	updatestatus "update-status/usecase/update_status"
)

var (
	quoteCodeDBNameExist = map[string][]configs.QuoteCodes{
		"AIG": {
			{
				QKbn:     "E",
				Sndc:     "T",
				LogicID:  "AIG0",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
		"CX2": {
			{
				QKbn:     "E",
				Sndc:     "M",
				LogicID:  "CX20",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
	}
	quoteCodeDBNameNotExist = map[string][]configs.QuoteCodes{
		"AIG": {
			{
				QKbn:     "E",
				Sndc:     "TLN",
				LogicID:  "AIG0",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
		"CX2": {
			{
				QKbn:     "E",
				Sndc:     "MT",
				LogicID:  "CX20",
				Endpoint: "127.0.0.1",
				DBName:   "tick",
			},
		},
	}
)

type argsDBName struct {
	updateService updatestatus.IUpdater
	quoteCode     map[string][]configs.QuoteCodes
	tick          configs.ArrayDatabaseStatus
	keihai        configs.ArrayDatabaseStatus
}

type expectDBName struct {
	groupStatus configs.GroupDatabaseStatusDefinition
	err         error
}

func Test_SetNewStatus_Database_Name(t *testing.T) {
	tests := []struct {
		name   string
		args   argsDBName
		expect expectDBName
	}{
		{
			name: "kei 1, data tick, have new status",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "1",
						Kei:      "1",
						Request: model.UpdateTypeDBName{
							DBName:    "tick",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  false,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  false,
							TheSecondKeiStatus: true,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "kei 1, data keihai, have new status",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "1",
						Request: model.UpdateTypeDBName{
							DBName:    "tick",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  false,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  false,
							TheSecondKeiStatus: true,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "kei 2, data tick, have new status",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "1",
						Kei:      "2",
						Request: model.UpdateTypeDBName{
							DBName:    "tick",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: false,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: false,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "kei 2, data keihai, have new status",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeDBName{
							DBName:    "tick",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: false,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: false,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "have new status, but not exists quoteCode in QuoteCodeDefinition ",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeDBName{
							DBName:    "tick",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameNotExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "not new status",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeDBName{
							DBName:    "tick",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameNotExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: false,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: false,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{
					Tick: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
					Kehai: configs.ArrayDatabaseStatus{
						{
							QKbn:               "E",
							Sndc:               "T",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: false,
						},
						{
							QKbn:               "E",
							Sndc:               "M",
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: false,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "missing request",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request:  nil,
					}, nil),
				quoteCode: quoteCodeDBNameNotExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: false,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: false,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{},
				err:         fmt.Errorf("invalid request update status"),
			},
		},
		{
			name: "wrong DBName",
			args: argsDBName{
				updateService: updatestatus.NewUpdateTypeDBNameService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeDBName{
							DBName:    "admin",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeDBNameNotExist,
				tick: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: true,
					},
				},
				keihai: configs.ArrayDatabaseStatus{
					{
						QKbn:               "E",
						Sndc:               "T",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: false,
					},
					{
						QKbn:               "E",
						Sndc:               "M",
						TheFirstKeiStatus:  true,
						TheSecondKeiStatus: false,
					},
				},
			},
			expect: expectDBName{
				groupStatus: configs.GroupDatabaseStatusDefinition{},
				err:         fmt.Errorf("invalid db name"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model.QuoteCodeDefinition = test.args.quoteCode
			model.TickDatabaseStatuses = test.args.tick
			model.KehaiDatabaseStatuses = test.args.keihai
			group, err := test.args.updateService.SetNewStatus()

			actual := expectDBName{
				err: err,
			}

			if group != nil {
				actual.groupStatus = *group
			}

			assert.Equal(t, test.expect, actual)
		})
	}
}
