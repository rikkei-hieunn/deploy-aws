/*
Package logger contain util log
*/
package logger

import "github.com/rs/zerolog/log"

// ShowLog show all log to console
func ShowLog(errors []error) {
	for index := range errors {
		log.Error().Msg(errors[index].Error())
	}
}
