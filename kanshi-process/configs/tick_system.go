package configs

import "errors"

// TickSystem structure of config about system
type TickSystem struct {
	RequestKinou     string `mapstructure:"TK_SYSTEM_REQUEST_KINOU"`
	RequestKanriID   string `mapstructure:"TK_SYSTEM_REQUEST_KANRI_ID"`
	RequestUserID    string `mapstructure:"TK_SYSTEM_REQUEST_USER_ID"`
	RequestSyubetu   string `mapstructure:"TK_SYSTEM_REQUEST_SYUBETU"`
	RequestQuoteCode string `mapstructure:"TK_SYSTEM_REQUEST_QUOTE_CODE"`
	RequestFromDate  string `mapstructure:"TK_SYSTEM_REQUEST_FROM_DATE"`
	RequestToDate    string `mapstructure:"TK_SYSTEM_REQUEST_TO_DATE"`
	RequestFromTime  string `mapstructure:"TK_SYSTEM_REQUEST_FROM_TIME"`
	RequestToTime    string `mapstructure:"TK_SYSTEM_REQUEST_TO_TIME"`
	RequestYobi      string `mapstructure:"TK_SYSTEM_REQUEST_YOBI"`
	RequestFunasi    string `mapstructure:"TK_SYSTEM_REQUEST_FUNASI"`
	RequestKikan     string `mapstructure:"TK_SYSTEM_REQUEST_KIKAN"`
	RequestElements  string `mapstructure:"TK_SYSTEM_REQUEST_ELEMENTS"`
}

// Request request to send toiawase
type Request []string

// Validate validate config
func (c *TickSystem) Validate() error {
	if len(c.RequestKinou) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_KINOU required")
	}
	if len(c.RequestKanriID) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_KANRI_ID required")
	}
	if len(c.RequestUserID) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_USER_ID required")
	}
	if len(c.RequestSyubetu) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_SYUBETU required")
	}
	if len(c.RequestQuoteCode) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_QUOTE_CODE required")
	}
	if len(c.RequestFromDate) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_FROM_DATE required")
	}
	if len(c.RequestToDate) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_TO_DATE required")
	}
	if len(c.RequestFunasi) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_FUNASI required")
	}
	if len(c.RequestKikan) == 0 {
		return errors.New("system TK_SYSTEM_REQUEST_KIKAN required")
	}

	return nil
}
