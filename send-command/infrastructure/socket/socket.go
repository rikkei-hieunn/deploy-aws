/*
Package socket implements logics about socket.
*/
package socket

import (
	"bufio"
	"fmt"
	"net"
	"send-command/model"
)

type socketHandler struct {
	writer *bufio.Writer
}

// NewSocketHandler init new socket handler
func NewSocketHandler() ISocketHandler {
	return &socketHandler{}
}

// InitConnection init socket connection and create writer for send socket message
func (s *socketHandler) InitConnection(host string, port int) error {
	connectionString := fmt.Sprintf("%s:%d", host, port)
	connection, err := net.Dial(model.ConnectionProtocol, connectionString)
	if err != nil {
		return err
	}
	s.writer = bufio.NewWriter(connection)

	return nil
}

// Send message
func (s *socketHandler) Send(command string) error {
	totalSent, err := s.writer.WriteString(command)
	if err != nil {
		return err
	}
	if totalSent != len(command) {
		return fmt.Errorf("send command fail")
	}

	return s.writer.Flush()
}
