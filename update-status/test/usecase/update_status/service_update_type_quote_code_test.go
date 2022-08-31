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
	quoteCodeExist = map[string][]configs.QuoteCodes{
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
	quoteCodeNotExist = map[string][]configs.QuoteCodes{
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

type argsQuoteCode struct {
	updateService updatestatus.IUpdater
	quoteCode     map[string][]configs.QuoteCodes
	tick          configs.ArrayDatabaseStatus
	keihai        configs.ArrayDatabaseStatus
}

type expectQuoteCote struct {
	groupStatus configs.GroupDatabaseStatusDefinition
	err         error
}

func Test_SetNewStatus_Quote_Code(t *testing.T) {
	tests := []struct {
		name   string
		args   argsQuoteCode
		expect expectQuoteCote
	}{
		{
			name: "kei 1, data tick, have new status",
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "1",
						Kei:      "1",
						Request: model.UpdateTypeQuoteCode{
							Kubun:     "E",
							Hassin:    "T",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeExist,
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
			expect: expectQuoteCote{
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
			name: "kei 1, data keihai, have new status",
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "1",
						Request: model.UpdateTypeQuoteCode{
							Kubun:     "E",
							Hassin:    "T",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeExist,
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
			expect: expectQuoteCote{
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
							TheFirstKeiStatus:  true,
							TheSecondKeiStatus: true,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "kei 2, data tick, have new status",
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "1",
						Kei:      "2",
						Request: model.UpdateTypeQuoteCode{
							Kubun:     "E",
							Hassin:    "M",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeExist,
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
			expect: expectQuoteCote{
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
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeQuoteCode{
							Kubun:     "E",
							Hassin:    "T",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeExist,
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
			expect: expectQuoteCote{
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
							TheSecondKeiStatus: true,
						},
					},
				},
				err: nil,
			},
		},
		{
			name: "not new status",
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeQuoteCode{
							Kubun:     "E",
							Hassin:    "T",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeNotExist,
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
			expect: expectQuoteCote{
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
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request:  nil,
					}, nil),
				quoteCode: quoteCodeNotExist,
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
			expect: expectQuoteCote{
				groupStatus: configs.GroupDatabaseStatusDefinition{},
				err:         fmt.Errorf("invalid request update status"),
			},
		},
		{
			name: "wrong DBName",
			args: argsQuoteCode{
				updateService: updatestatus.NewUpdateTypeQuoteCodeService(
					&configs.TickSystem{
						DataType: "2",
						Kei:      "2",
						Request: model.UpdateTypeQuoteCode{
							Kubun:     "F",
							Hassin:    "X",
							NewStatus: false,
						},
					}, nil),
				quoteCode: quoteCodeNotExist,
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
			expect: expectQuoteCote{
				groupStatus: configs.GroupDatabaseStatusDefinition{},
				err:         fmt.Errorf("invalid kubun and hassin"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model.QuoteCodeDefinition = test.args.quoteCode
			model.TickDatabaseStatuses = test.args.tick
			model.KehaiDatabaseStatuses = test.args.keihai
			group, err := test.args.updateService.SetNewStatus()

			actual := expectQuoteCote{
				err: err,
			}

			if group != nil {
				actual.groupStatus = *group
			}

			assert.Equal(t, test.expect, actual)
		})
	}
}
