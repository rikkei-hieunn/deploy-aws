/*
Package repository implements logics repository.
*/
package repository

import (
	"bytes"
)

// IS3Repository interface S3 repository
type IS3Repository interface {
	Download(path string) ([]byte, error)
	Upload(prefix string, data bytes.Buffer) error
}

// ISocketRepository Structure of interface socket repository
type ISocketRepository interface {
	InitWriter(machineName string, port int) error
	SendCommand(command string) error
}
