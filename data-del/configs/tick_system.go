/*
Package configs contains configuration.
*/
package configs

import "fmt"

//TickSystem tick system define all configuration in program
type TickSystem struct {
	CommonDefinitionObject           string `mapstructure:"TK_SYSTEM_SHARE_INFORMATION_OBJECT"`
	ExpireDefinitionObject           string `mapstructure:"TK_SYSTEM_EXPIRE_INFORMATION_OBJECT"`
	TickDB1EndpointDefinitionObject  string `mapstructure:"TK_SYSTEM_TICK_DB1_ENDPOINT_INFORMATION_OBJECT"`
	TickDB2EndpointDefinitionObject  string `mapstructure:"TK_SYSTEM_TICK_DB2_ENDPOINT_INFORMATION_OBJECT"`
	KehaiDB1EndpointDefinitionObject string `mapstructure:"TK_SYSTEM_KEHAI_DB1_ENDPOINT_INFORMATION_OBJECT"`
	KehaiDB2EndpointDefinitionObject string `mapstructure:"TK_SYSTEM_KEHAI_DB2_ENDPOINT_INFORMATION_OBJECT"`
	DevelopEnvironment               bool   `mapstructure:"TK_SYSTEM_DEVELOP_ENVIRONMENT"`

	S3Region            string
	S3Bucket            string
	NumberOfDeletedDays int
	ExpiredDays         []TickExpire
	ExpiredDaysAll      TickExpire
	TablePrefix         map[int]string
}

//TickExpire struct to map with config
type TickExpire struct {
	QKBN   string `json:"QKBN"`
	SNDC   string `json:"SNDC"`
	Expire int    `json:"Expired"`
}

//TickTablePrefix struct to map with config
type TickTablePrefix struct {
	DataType int    `mapstructure:"DataType"`
	Prefix   string `mapstructure:"Prefix"`
}

// Validate validate log config
func (c *TickSystem) Validate() error {
	if len(c.CommonDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_SHARE_INFORMATION_OBJECT required")
	}
	if len(c.ExpireDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_EXPIRE_INFORMATION_OBJECT required")
	}
	if len(c.TickDB1EndpointDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_TICK_DB1_ENDPOINT_INFORMATION_OBJECT required")
	}
	if len(c.TickDB2EndpointDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_TICK_DB2_ENDPOINT_INFORMATION_OBJECT required")
	}
	if len(c.KehaiDB1EndpointDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_KEHAI_DB1_ENDPOINT_INFORMATION_OBJECT required")
	}
	if len(c.KehaiDB2EndpointDefinitionObject) == 0 {
		return fmt.Errorf("system TK_SYSTEM_KEHAI_DB2_ENDPOINT_INFORMATION_OBJECT required")
	}

	return nil
}
