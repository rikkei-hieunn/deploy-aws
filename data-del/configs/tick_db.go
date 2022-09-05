/*
Package configs contains configuration.
*/
package configs

import "fmt"

// TickDB contains sorting DB configuration
type TickDB struct {
	Port              int    `json:"TK_DB_PORT"`
	User              string `json:"TK_DB_USER"`
	Password          string `json:"TK_DB_PASSWORD"`
	MaxOpenConnection int    `json:"TK_DB_MAX_OPEN_CONNECTION"`
	MaxIdleConnection int    `json:"TK_DB_MAX_IDLE_CONNECTION"`
	DriverName        string `json:"TK_DB_DRIVER_NAME"`
	RetryTimes        int    `json:"TK_DB_RETRY_TIMES"`
	RetryWaitTimes    int    `json:"TK_DB_RETRY_WAIT_MS"`

	Endpoints map[string]map[string][]string
}

//EndPoint  structure of endpoint
type EndPoint struct {
	TKQKBN     string `json:"QKBN"`
	SNDC       string `json:"SNDC"`
	DBEndpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName     string `json:"TKDBNAME"`
}

// Validate validate log config
func (c *TickDB) Validate() error {
	if c.Port == 0 {
		return fmt.Errorf("database TK_DB_PORT required")
	}
	if c.Port < 0 {
		return fmt.Errorf("invalid TK_DB_PORT")
	}
	if len(c.User) == 0 {
		return fmt.Errorf("database TK_DB_USER required")
	}
	if len(c.Password) == 0 {
		return fmt.Errorf("database TK_DB_PASSWORD required")
	}
	if c.MaxOpenConnection == 0 {
		return fmt.Errorf("database TK_DB_MAX_OPEN_CONNECTION required")
	}
	if c.MaxOpenConnection < 0 {
		return fmt.Errorf("invalid TK_DB_MAX_OPEN_CONNECTION")
	}
	if c.MaxIdleConnection == 0 {
		return fmt.Errorf("database TK_DB_MAX_IDLE_CONNECTION required")
	}
	if c.MaxIdleConnection < 0 {
		return fmt.Errorf("invalid TK_DB_MAX_IDLE_CONNECTION")
	}
	if len(c.DriverName) == 0 {
		return fmt.Errorf("database TK_DB_DRIVER_NAME required")
	}

	return nil
}
