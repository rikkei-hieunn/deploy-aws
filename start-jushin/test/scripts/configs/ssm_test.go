package configs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"start-jushin/configs"
	"testing"
)

func Test_SSM_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.Group
		expect error
	}{
		{name: "miss logic group id ",
			args: configs.Group{
				GroupID:       "",
				TopicName:     "topicname",
				Types:         "aaa",
				CommandPort:   "2206",
				TickHostName:  "host",
				KehaiHostName: "host",
				GroupLine:     "1",
			},
			expect: fmt.Errorf("TK_LOGIC_GROUP is required "),
		},
		{name: " missing topic name",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "",
				Types:         "aaa",
				CommandPort:   "2206",
				TickHostName:  "host",
				KehaiHostName: "host",
				GroupLine:     "1",
			},
			expect: fmt.Errorf("TK_TOPIC_NAME is required ")},
		{name: " missing type ",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "aaa",
				Types:         "",
				CommandPort:   "2206",
				TickHostName:  "host",
				KehaiHostName: "host",
				GroupLine:     "1",
			},
			expect: fmt.Errorf("TK_TYPES is required ")},
		{name: " missing command port ",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "aaa",
				Types:         "aa",
				CommandPort:   "",
				TickHostName:  "host",
				KehaiHostName: "host",
				GroupLine:     "1",
			},
			expect: fmt.Errorf("TK_COMMAND_PORT is required ")},
		{name: " missing tick host name ",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "aaa",
				Types:         "aaa",
				CommandPort:   "3333",
				TickHostName:  "",
				KehaiHostName: "host",
				GroupLine:     "1",
			},
			expect: fmt.Errorf("TK_TICK_HOST_NAME is required ")},
		{name: " missing kehai host name ",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "aaa",
				Types:         "aaa",
				CommandPort:   "3333",
				TickHostName:  "aaa",
				KehaiHostName: "",
				GroupLine:     "1",
			},
			expect: fmt.Errorf("TK_KEHAI_HOST_NAME is required ")},
		{name: " missing group line ",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "aaa",
				Types:         "aaa",
				CommandPort:   "3333",
				TickHostName:  "aaa",
				KehaiHostName: "host",
				GroupLine:     "",
			},
			expect: fmt.Errorf("TK_GROUP_LINE is required ")},

		{name: "Validate success",
			args: configs.Group{
				GroupID:       "1",
				TopicName:     "topicname",
				Types:         "aaa",
				CommandPort:   "2206",
				TickHostName:  "host",
				KehaiHostName: "host",
				GroupLine:     "1",
			},
			expect: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			if tt.name != "Validate success" {
				assert.Equal(t, result.Error(), tt.expect.Error())
			} else {
				assert.Equal(t, result, nil)
			}
		})
	}
}
