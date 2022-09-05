// Package filebus define filebus interface
package filebus

// IFilebusHandler define function of Filebus handler
type IFilebusHandler interface {
	Download(path, file string) ([]byte, error)
}
