/*
Package configs contains configuration.
*/
package configs

import (
	"errors"
)

// TickFileBus struct contain info file bus
type TickFileBus struct {
	Username       string `json:"TK_FILEBUS_USER"`
	Password       string `json:"TK_FILEBUS_PASSWORD"`
	Hostname       string `json:"TK_FILEBUS_HOST"`
	Port           int    `json:"TK_FILEBUS_PORT"`
	URLDownload    string `json:"TK_FILEBUS_URL_DOWNLOAD_FILE"`
	PathCalendar1  string `json:"TK_FILEBUS_PATH_CALENDAR1_FILE"`
	PathCalendar2  string `json:"TK_FILEBUS_PATH_CALENDAR2_FILE"`
	RetryTimes     int    `json:"TK_FILEBUS_RETRY_TIMES"`
	RetryWaitTimes int    `json:"TK_FILEBUS_RETRY_WAIT"`
}

// Validate validate Filebus config
func (s *TickFileBus) Validate() error {
	if len(s.Username) == 0 {
		return errors.New("filebus TK_FILEBUS_USER required")
	}
	if len(s.Password) == 0 {
		return errors.New("filebus TK_FILEBUS_PASSWORD required")
	}
	if len(s.Hostname) == 0 {
		return errors.New("filebus TK_FILEBUS_HOSTNAME required")
	}
	if len(s.URLDownload) == 0 {
		return errors.New("filebus TK_FILEBUS_URL_DOWNLOAD_FILE required")
	}
	if len(s.PathCalendar1) == 0 {
		return errors.New("filebus TK_FILEBUS_PATH_CALENDAR1_FILE required")
	}
	if len(s.PathCalendar2) == 0 {
		return errors.New("filebus TK_FILEBUS_PATH_CALENDAR2_FILE required")
	}
	if s.Port == 0 {
		return errors.New("filebus TK_FILEBUS_PORT required")
	}
	if s.Port < 0 {
		return errors.New("ìnvalid TK_FILEBUS_PORT")
	}
	if s.RetryTimes < 0 {
		return errors.New("ìnvalid TK_FILEBUS_RETRY_TIMES")
	}
	if s.RetryWaitTimes < 0 {
		return errors.New("ìnvalid TK_FILEBUS_RETRY_WAIT")
	}

	return nil
}
