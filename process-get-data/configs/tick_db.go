package configs

import "errors"

// TickDB structure of config about database
type TickDB struct {
	Port              uint   `json:"TK_DB_PORT"`
	User              string `json:"TK_DB_USER"`
	Password          string `json:"TK_DB_PASSWORD"`
	MaxOpenConnection int    `json:"TK_DB_MAX_OPEN_CONNECTION"`
	MaxIdleConnection int    `json:"TK_DB_MAX_IDLE_CONNECTION"`
	DriverName        string `json:"TK_DB_DRIVER_NAME"`
	RetryTimes        int    `json:"TK_DB_RETRY_TIMES"`
	RetryWaitMs       int    `json:"TK_DB_RETRY_WAIT_MS"`
}

// TableNamePrefix structure of config prefix of table
type TableNamePrefix []Prefix

// Prefix prefix table
type Prefix struct {
	DataType int    `mapstructure:"DataType" json:"DataType"`
	Prefix   string `mapstructure:"Prefix" json:"Prefix"`
}

// Validate validate log config
func (c *TickDB) Validate() error {
	if c.Port == 0 {
		return errors.New("TK_DB_PORT")
	}
	if len(c.User) == 0 {
		return errors.New("TK_DB_USER")
	}
	if len(c.Password) == 0 {
		return errors.New("TK_DB_PASSWORD")
	}
	if c.MaxOpenConnection == 0 {
		return errors.New("TK_DB_MAX_OPEN_CONNECTION")
	}
	if c.MaxIdleConnection == 0 {
		return errors.New("TK_DB_MAX_IDLE_CONNECTION")
	}
	if len(c.DriverName) == 0 {
		return errors.New("TK_DB_DRIVER_NAME")
	}

	return nil
}
