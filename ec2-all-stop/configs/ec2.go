/*
Package configs declare all configs use in application.
*/
package configs

// EC2 structure define list of instance's id
type EC2 struct {
	InstanceIds         []string
	HostNameDefinitions []HostNameDefinition
}

// HostNameDefinition structure define hostname
type HostNameDefinition struct {
	TickHostName  string `json:"TK_TICK_HOST_NAME"`
	KehaiHostName string `json:"TK_KEHAI_HOST_NAME"`
}
