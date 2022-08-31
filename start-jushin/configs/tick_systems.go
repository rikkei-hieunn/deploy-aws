package configs

import (
	"fmt"
	"start-jushin/model"
	"strings"
)

// TickSystem contains application configuration
type TickSystem struct {
	GroupDefinitionForFirstKei  string `mapstructure:"TK_SYSTEM_FIRST_KEI_S3_PATH"`
	GroupDefinitionForSecondKei string `mapstructure:"TK_SYSTEM_SECOND_KEI_S3_PATH"`
	InstancePathKey             string `mapstructure:"TK_SYSTEM_INSTANCE_PATH_ENV_KEY"`
	DevelopEnvironment          bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Bucket string
	S3Region string
}

// Validate func validates application configuration
func (s *TickSystem) Validate() error {
	if strings.TrimSpace(s.GroupDefinitionForFirstKei) == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_FIRST_KEI_S3_PATH is required ")
	}
	if strings.TrimSpace(s.GroupDefinitionForSecondKei) == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_SECOND_KEI_S3_PATH is required ")
	}
	if strings.TrimSpace(s.InstancePathKey) == model.EmptyString {
		return fmt.Errorf("TK_SYSTEM_INSTANCE_PATH_ENV_KEY is required ")
	}

	return nil
}
