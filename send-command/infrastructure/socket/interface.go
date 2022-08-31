/*
Package socket implements logics about socket.
*/
package socket

// ISocketHandler Structure of interface socket handler
type ISocketHandler interface {
	InitConnection(host string, port int) error
	Send(command string) error
}
