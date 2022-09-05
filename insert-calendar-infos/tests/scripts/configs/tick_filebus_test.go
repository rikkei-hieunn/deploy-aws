package configs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/model"
	"testing"
)

func Test_TickFileBus_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickFileBus
		expect error
	}{
		{
			name: "TK_FILEBUS_USER is empty",
			args: configs.TickFileBus{
				Port:           8089,
				Username:       model.EmptyString,
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_USER required"),
		},
		{
			name: "TK_FILEBUS_PASSWORD is empty",
			args: configs.TickFileBus{
				Port:           8089,
				Username:       "admin",
				Password:       model.EmptyString,
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_PASSWORD required"),
		},
		{
			name: "TK_FILEBUS_HOSTNAME is empty",
			args: configs.TickFileBus{
				Port:           8089,
				Username:       "admin",
				Password:       "admin",
				Hostname:       model.EmptyString,
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_HOSTNAME required"),
		},
		{
			name: "TK_FILEBUS_URL_DOWNLOAD_FILE is empty",
			args: configs.TickFileBus{
				Port:           8089,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    model.EmptyString,
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_URL_DOWNLOAD_FILE required"),
		},
		{
			name: "TK_FILEBUS_PATH_CALENDAR1_FILE is empty",
			args: configs.TickFileBus{
				Port:           8089,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  model.EmptyString,
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_PATH_CALENDAR1_FILE required"),
		},
		{
			name: "TK_FILEBUS_PATH_CALENDAR2_FILE is empty",
			args: configs.TickFileBus{
				Port:           8089,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  model.EmptyString,
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_PATH_CALENDAR2_FILE required"),
		},
		{
			name: "TK_FILEBUS_PORT is empty",
			args: configs.TickFileBus{
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("filebus TK_FILEBUS_PORT required"),
		},
		{
			name: "ìnvalid TK_FILEBUS_PORT",
			args: configs.TickFileBus{
				Port:           -1000,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("ìnvalid TK_FILEBUS_PORT"),
		},
		{
			name: "ìnvalid TK_FILEBUS_RETRY_TIMES",
			args: configs.TickFileBus{
				Port:           1000,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     -3,
				RetryWaitTimes: 3000,
			},
			expect: errors.New("ìnvalid TK_FILEBUS_RETRY_TIMES"),
		},
		{
			name: "ìnvalid TK_FILEBUS_RETRY_WAIT",
			args: configs.TickFileBus{
				Port:           1000,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: -3000,
			},
			expect: errors.New("ìnvalid TK_FILEBUS_RETRY_WAIT"),
		},
		{
			name: "validate success",
			args: configs.TickFileBus{
				Port:           1000,
				Username:       "admin",
				Password:       "admin",
				Hostname:       "10.1.40.178",
				URLDownload:    "/file/download",
				PathCalendar1:  "sample-file/",
				PathCalendar2:  "sample-file/",
				RetryTimes:     3,
				RetryWaitTimes: 3,
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			assert.Equal(t, result, tt.expect)
		})
	}
}
