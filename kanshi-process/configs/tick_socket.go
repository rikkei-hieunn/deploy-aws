package configs

import "errors"

// TickSocket TickSystem structure of config about system
type TickSocket struct {
	ConnectionType string         `mapstructure:"TK_SOCKET_CONNECTION_TYPE"`
	ToiawaseSocket ToiawaseSocket `mapstructure:"TK_SOCKET_TOIAWASE"`
	FuriwakeSocket FuriwakeSocket `mapstructure:"TK_SOCKET_FURIWAKE"`
}

// ToiawaseSocket struct define host and port to connection toiawase server
type ToiawaseSocket struct {
	Host  string `mapstructure:"TOIAWASE_HOST"`
	Ports []int  `mapstructure:"TOIAWASE_PORT"`
}

// FuriwakeSocket struct define host and port to connection furiwake server
type FuriwakeSocket struct {
	Host string `mapstructure:"FURIWAKE_HOST"`
	Port int    `mapstructure:"FURIWAKE_PORT"`
}

// Validate validate config socket
func (c *TickSocket) Validate() error {
	if len(c.ConnectionType) == 0 {
		return errors.New("system TK_SOCKET_CONNECTION_TYPE required")
	}
	if len(c.ToiawaseSocket.Ports) == 0 {
		return errors.New("system TOIAWASE_PORT required")
	}
	if len(c.ToiawaseSocket.Ports) != 3 {
		return errors.New("system TOIAWASE_PORT required")
	}

	for index := range c.ToiawaseSocket.Ports {
		if c.ToiawaseSocket.Ports[index] < 0 {
			return errors.New("invalid TOIAWASE_PORT")
		}
	}

	if len(c.ToiawaseSocket.Host) == 0 {
		return errors.New("system TOIAWASE_HOST required")
	}
	if len(c.FuriwakeSocket.Host) == 0 {
		return errors.New("system FURIWAKE_HOST required")
	}
	if c.FuriwakeSocket.Port == 0 {
		return errors.New("system FURIWAKE_PORT required")
	}
	if c.FuriwakeSocket.Port < 0 {
		return errors.New("invalid FURIWAKE_PORT")
	}

	return nil
}
