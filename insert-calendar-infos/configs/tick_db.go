/*
Package configs contains configuration.
*/
package configs

import (
	"errors"
)

// TickDB contains sorting DB configuration
type TickDB struct {
	Port              int    `json:"TK_DB_PORT"`
	User              string `json:"TK_DB_USER"`
	Password          string `json:"TK_DB_PASSWORD"`
	MaxOpenConnection int    `json:"TK_DB_MAX_OPEN_CONNECTION"`
	MaxIdleConnection int    `json:"TK_DB_MAX_IDLE_CONNECTION"`
	DriverName        string `json:"TK_DB_DRIVER_NAME"`
	RetryTimes        int    `json:"TK_DB_RETRY_TIMES"`
	RetryWaitMs       int    `json:"TK_DB_RETRY_WAIT_MS"`

	CalendarEndpoint CalendarInfo
}

// CalendarInfo contain information table calendar table
type CalendarInfo struct {
	TableName      string `json:"TK_CALENDAR_TABLE_NAME"`
	DBKei1Endpoint string `json:"TK_CALENDAR_KEI1_ENDPOINT"`
	DBKei1Name     string `json:"TK_CALENDAR_KEI1_DBNAME"`
	DBKei2Endpoint string `json:"TK_CALENDAR_KEI2_ENDPOINT"`
	DBKei2Name     string `json:"TK_CALENDAR_KEI2_DBNAME"`
}

// Validate validate Calendar info config
func (c *CalendarInfo) Validate() error {
	if len(c.TableName) == 0 {
		return errors.New("database TK_CALENDAR_TABLE_NAME required")
	}
	if len(c.DBKei1Endpoint) == 0 {
		return errors.New("database TK_CALENDAR_KEI1_ENDPOINT required")
	}
	if len(c.DBKei1Name) == 0 {
		return errors.New("database TK_CALENDAR_KEI1_DBNAME required")
	}
	if len(c.DBKei2Endpoint) == 0 {
		return errors.New("database TK_CALENDAR_KEI2_ENDPOINT required")
	}
	if len(c.DBKei2Name) == 0 {
		return errors.New("database TK_CALENDAR_KEI2_DBNAME required")
	}

	return nil
}

// Validate validate Tick DB config
func (c *TickDB) Validate() error {
	if c.Port == 0 {
		return errors.New("database TK_DB_PORT required")
	}
	if c.Port < 0 {
		return errors.New("invalid TK_DB_PORT")
	}
	if len(c.User) == 0 {
		return errors.New("database TK_DB_USER required")
	}
	if len(c.Password) == 0 {
		return errors.New("database TK_DB_PASSWORD required")
	}
	if c.MaxOpenConnection == 0 {
		return errors.New("database TK_DB_MAX_OPEN_CONNECTION required")
	}
	if c.MaxOpenConnection < 0 {
		return errors.New("invalid TK_DB_MAX_OPEN_CONNECTION")
	}
	if c.MaxIdleConnection == 0 {
		return errors.New("database TK_DB_MAX_IDLE_CONNECTION required")
	}
	if c.MaxIdleConnection < 0 {
		return errors.New("invalid TK_DB_MAX_IDLE_CONNECTION")
	}
	if len(c.DriverName) == 0 {
		return errors.New("database TK_DB_DRIVER_NAME required")
	}

	return nil
}
