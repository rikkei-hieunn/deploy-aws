package model

// RequestAll type of request for all
type RequestAll struct {
	Kei     string
	Command string
}

// RequestGroupID type of request for group ID
type RequestGroupID struct {
	Kei      string
	DataType string
	GroupID  string
	Command  string
}

// RequestLine type of request for group line
type RequestLine struct {
	Kei       string
	GroupLine string
	Command   string
}

// RequestToiawase type of request for toiswase
type RequestToiawase struct {
	Command string
}

// GroupDefinition structure define group
type GroupDefinition struct {
	GroupID          string `json:"TKLOGIC_GROUP"`
	Types            string `json:"TKTYPES"`
	Port             string `json:"TK_COMMAND_PORT"`
	TickMachineName  string `json:"TK_TICK_HOST_NAME"`
	KehaiMachineName string `json:"TK_KEHAI_HOST_NAME"`
	Line             string `json:"TK_GROUP_LINE"`
}
