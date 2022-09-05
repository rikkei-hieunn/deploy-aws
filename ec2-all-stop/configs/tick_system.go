package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	GroupsDefinitionKei1Object string `mapstructure:"TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT"`
	GroupsDefinitionKei2Object string `mapstructure:"TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT"`
	DevelopEnvironment         bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	Kei      string
	S3Region string
	S3Bucket string
}

// Validate validate object tick system
func (c *TickSystem) Validate() error {
	if len(c.GroupsDefinitionKei1Object) == 0 {
		return errors.New("system TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT required")
	}
	if len(c.GroupsDefinitionKei2Object) == 0 {
		return errors.New("system TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT required")
	}

	return nil
}
