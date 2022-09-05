/*
Package configs contains configuration.
*/
package configs

import "fmt"

// TickDB contains sorting DB configuration
type TickDB struct {
	Port              uint   `json:"TK_DB_PORT"`
	User              string `json:"TK_DB_USER"`
	Password          string `json:"TK_DB_PASSWORD"`
	MaxOpenConnection int    `json:"TK_DB_MAX_OPEN_CONNECTION"`
	MaxIdleConnection int    `json:"TK_DB_MAX_IDLE_CONNECTION"`
	DriverName        string `json:"TK_DB_DRIVER_NAME"`
	RetryTimes        int    `json:"TK_DB_RETRY_TIMES"`
	RetryWaitTimes    int    `json:"TK_DB_RETRY_WAIT_MS"`

	Endpoints map[string][]string
}

// EndPoint contains DB configuration host and DB name
type EndPoint struct {
	TKQKBN     string `json:"QKBN"`
	SNDC       string `json:"SNDC"`
	DBEndpoint string `json:"TKDB_MASTER_ENDPOINT"`
	DBName     string `json:"TKDBNAME"`
}

// NumberPrefix contains prefix number quote code
type NumberPrefix struct {
	Start string `json:"TKSTART_PROPERNAME"`
	End   string `json:"TKEND_PROPERNAME"`
}

// CandleManagement contains table name
type CandleManagement struct {
	TableName string `json:"TK_CANDLE_MANAGEMENT_TABLE_NAME"`
}

// Validate validate log config
func (c *TickDB) Validate() error {
	if c.Port == 0 {
		return fmt.Errorf("database TK_DB_PORT required")
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
	if c.MaxIdleConnection == 0 {
		return fmt.Errorf("database TK_DB_MAX_IDLE_CONNECTION required")
	}
	if len(c.DriverName) == 0 {
		return fmt.Errorf("database TK_DB_DRIVER_NAME required")
	}

	return nil
}
