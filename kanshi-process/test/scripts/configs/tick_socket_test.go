package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"kanshi-process/configs"
	"kanshi-process/model"
	"testing"
)

func Test_TickSocket_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSocket
		expect error
	}{
		{
			name: "TK_SOCKET_CONNECTION_TYPE is empty",
			args: configs.TickSocket{
				ConnectionType: model.EmptyString,
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
						9002,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: 9000,
				},
			},
			expect: errors.New("system TK_SOCKET_CONNECTION_TYPE required"),
		},
		{
			name: "Toiawase ports are empty",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host:  "localhost",
					Ports: []int{},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: 9000,
				},
			},
			expect: errors.New("system TOIAWASE_PORT required"),
		},
		{
			name: "Toiawase ports does not equal 3",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: 9000,
				},
			},
			expect: errors.New("system TOIAWASE_PORT required"),
		},
		{
			name: "Toiawase ports invalid",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
						-9001,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: 9000,
				},
			},
			expect: errors.New("invalid TOIAWASE_PORT"),
		},
		{
			name: "TOIAWASE_HOST is empty",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: model.EmptyString,
					Ports: []int{
						9000,
						9001,
						9002,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: 9000,
				},
			},
			expect: errors.New("system TOIAWASE_HOST required"),
		},
		{
			name: "FURIWAKE_HOST is empty",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
						9002,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: model.EmptyString,
					Port: 9000,
				},
			},
			expect: errors.New("system FURIWAKE_HOST required"),
		},
		{
			name: "Furiwake port is empty",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
						9002,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
				},
			},
			expect: errors.New("system FURIWAKE_PORT required"),
		},
		{
			name: "Furiwake port invalid",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
						9002,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: -1000,
				},
			},
			expect: errors.New("invalid FURIWAKE_PORT"),
		},
		{
			name: "Validate success",
			args: configs.TickSocket{
				ConnectionType: "tcp",
				ToiawaseSocket: configs.ToiawaseSocket{
					Host: "localhost",
					Ports: []int{
						9000,
						9001,
						9002,
					},
				},
				FuriwakeSocket: configs.FuriwakeSocket{
					Host: "localhost",
					Port: 1000,
				},
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
