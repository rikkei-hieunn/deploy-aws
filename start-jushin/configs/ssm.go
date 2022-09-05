package configs

import (
	"fmt"
	"start-jushin/model"
	"strings"
)

// SSM struct
type SSM struct {
	InstanceID string
	Commands   []string
	MaxWaitTime   int `mapstructure:"MAX_WAIT_TIME"`
	RetryWaitTime int `mapstructure:"RETRY_WAIT_TIME"`
}

// Group struct
type Group struct {
	GroupID   string `json:"TK_LOGIC_GROUP"`
	TopicName string `json:"TK_TOPIC_NAME"`
	Types         string `json:"TK_TYPES"`
	CommandPort   string `json:"TK_COMMAND_PORT"`
	TickHostName  string `json:"TK_TICK_HOST_NAME"`
	KehaiHostName string `json:"TK_KEHAI_HOST_NAME"`
	GroupLine     string `json:"TK_GROUP_LINE"`
}

// Validate func validate application configuration
func (s *Group) Validate() error {
	if strings.TrimSpace(s.GroupID) == model.EmptyString {
		return fmt.Errorf("TK_LOGIC_GROUP is required ")
	}
	if strings.TrimSpace(s.TopicName) == model.EmptyString {
		return fmt.Errorf("TK_TOPIC_NAME is required ")
	}
	if strings.TrimSpace(s.Types) == model.EmptyString {
		return fmt.Errorf("TK_TYPES is required ")
	}
	if strings.TrimSpace(s.CommandPort) == model.EmptyString {
		return fmt.Errorf("TK_COMMAND_PORT is required ")
	}
	if strings.TrimSpace(s.TickHostName) == model.EmptyString {
		return fmt.Errorf("TK_TICK_HOST_NAME is required ")
	}
	if strings.TrimSpace(s.KehaiHostName) == model.EmptyString {
		return fmt.Errorf("TK_KEHAI_HOST_NAME is required ")
	}
	if strings.TrimSpace(s.GroupLine) == model.EmptyString {
		return fmt.Errorf("TK_GROUP_LINE is required ")
	}

	return nil
}
