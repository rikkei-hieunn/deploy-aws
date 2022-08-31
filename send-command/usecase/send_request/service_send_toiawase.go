package sendrequest

import (
	"fmt"
	"os"
	"send-command/configs"
	"send-command/infrastructure/repository"
	"send-command/model"
	"strconv"
)

type serviceSendRequestToiawase struct {
	config           *configs.TickSystem
	socketRepository repository.ISocketRepository
}

// NewSendToiawaseService constructor init service send command
func NewSendToiawaseService(cfg *configs.TickSystem, socketRepo repository.ISocketRepository) ISender {
	return &serviceSendRequestToiawase{
		config:           cfg,
		socketRepository: socketRepo,
	}
}

// HandleRequest handle request and send command
func (s *serviceSendRequestToiawase) HandleRequest() error {
	request, ok := s.config.Request.(model.RequestToiawase)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	// get toiawase host from os environment
	toiawaseHost := os.Getenv(model.KeyToiawaseHost)
	if toiawaseHost == model.EmptyString {
		return fmt.Errorf("invalid toiawase host")
	}

	// get toiawase the first port from os environment
	toiawaseFirstPort := os.Getenv(model.KeyToiawaseTheFirstPort)
	if toiawaseFirstPort == model.EmptyString {
		return fmt.Errorf("invalid toiawase the first port")
	}
	firstPortNumber, err := strconv.Atoi(toiawaseFirstPort)
	if err != nil {
		return fmt.Errorf("invalid toiawase the first port")
	}

	// get toiawase the second port from os environment
	toiawaseSecondPort := os.Getenv(model.KeyToiawaseTheSecondPort)
	if toiawaseSecondPort == model.EmptyString {
		return fmt.Errorf("invalid toiawase the second port")
	}
	secondPortNumber, err := strconv.Atoi(toiawaseSecondPort)
	if err != nil {
		return fmt.Errorf("invalid toiawase the second port")
	}

	// get toiawase the third port from os environment
	toiawaseThirdPort := os.Getenv(model.KeyToiawaseTheThirdPort)
	if toiawaseThirdPort == model.EmptyString {
		return fmt.Errorf("invalid toiawase the third port")
	}
	thirdPortNumber, err := strconv.Atoi(toiawaseThirdPort)
	if err != nil {
		return fmt.Errorf("invalid toiawase the third port")
	}

	// send command to the first port
	err = s.socketRepository.InitWriter(toiawaseHost, firstPortNumber)
	if err != nil {
		return err
	}

	err = s.socketRepository.SendCommand(request.Command)
	if err != nil {
		return err
	}

	// send command to the second port
	err = s.socketRepository.InitWriter(toiawaseHost, secondPortNumber)
	if err != nil {
		return err
	}

	err = s.socketRepository.SendCommand(request.Command)
	if err != nil {
		return err
	}

	// send command to the third port
	err = s.socketRepository.InitWriter(toiawaseHost, thirdPortNumber)
	if err != nil {
		return err
	}

	err = s.socketRepository.SendCommand(request.Command)
	if err != nil {
		return err
	}

	return nil
}
