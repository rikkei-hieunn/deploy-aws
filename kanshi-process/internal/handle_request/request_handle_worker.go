// Package handlerequest for handle request worker
package handlerequest

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"kanshi-process/configs"
	"kanshi-process/model"
	"net"
	"strconv"
	"strings"
	"time"
)

type worker struct {
	config  *configs.TickSocket
	request configs.Request
}

// NewWorker constructor init new worker
func NewWorker(request configs.Request, cfg *configs.TickSocket) IWorker {
	return &worker{
		request: request,
		config:  cfg,
	}
}

// Process process send request
func (w *worker) Process(ctx context.Context, host string, port int) {
	timeSleepConnect := model.DefaultTimeSleepConnect
	retryTimesConnect := model.DefaultRetryTimesConnect

	var err error
	var totalError int
	var connection net.Conn

	defer func() {
		if connection != nil {
			_ = connection.Close()
		}
	}()

	for {
		// connect to socket server
		for index := 1; index <= retryTimesConnect; index++ {
			connectionString := fmt.Sprintf("%s:%d", host, port)
			connection, err = net.Dial(w.config.ConnectionType, connectionString)
			if err != nil || connection == nil {
				log.Error().Msgf("connection fail, host: %s, port: %d", host, port)
				time.Sleep(time.Duration(timeSleepConnect) * time.Second)

				continue
			}

			break
		}

		if connection == nil {
			// TODO push log to Senjiu system here
			timeSleepConnect = model.ErrorTimeSleepConnect
			retryTimesConnect = model.ErrorRetryTimesConnect

			continue
		}

		// send request to socket server
		writer := bufio.NewWriter(connection)
		reader := bufio.NewReader(connection)
		for index := range w.request {
			_, err = writer.WriteString(w.request[index])
			if err != nil {
				log.Error().Msgf("send request fail, request: %s", strings.Join(w.request, model.EmptyString))

				continue
			}
			err = writer.Flush()
			if err != nil {
				log.Error().Msgf("send request fail, request: %s", strings.Join(w.request, model.EmptyString))

				continue
			}
		}

		// receive response data
		for {
			var startHeader bytes.Buffer
			totalHeaderLength := model.TotalResponseBytesForFirstType
			if w.request[0] == model.SecondFunctionType {
				totalHeaderLength = model.TotalResponseBytesForSecondType
			}

			// receive header
			isSuccess := true
			for totalHeaderLength > 0 {
				receiverByte, err := reader.ReadByte()
				if err != nil {
					log.Error().Msg("receive response fail")
					isSuccess = false

					break
				}
				startHeader.WriteByte(receiverByte)
				totalHeaderLength--

				if startHeader.Len() == 2 && isSuccess {
					// Get kinou response
					if startHeader.String()[:2] == model.ErrorFunctionType {
						totalHeaderLength = model.TotalResponseBytesForError
						isSuccess = false
					}
				}
			}

			var compressData bytes.Buffer
			var totalCompressDataLength int
			if !isSuccess {
				goto handle_error
			}

			if startHeader.Bytes()[94] == '9' {
				totalError = 0
				timeSleepConnect = model.DefaultTimeSleepConnect
				retryTimesConnect = model.DefaultRetryTimesConnect

				break
			}

			// get total response data length
			totalCompressDataLength, err = strconv.Atoi(strings.TrimSpace(startHeader.String()[len(startHeader.String())-8:]))
			if err != nil {
				goto handle_error
			}

			// Receive compress data
			for totalCompressDataLength > 0 {
				receiverByte, err := reader.ReadByte()
				if err != nil {
					goto handle_error
				}
				compressData.WriteByte(receiverByte)
				totalCompressDataLength--
			}
			totalError = 0
			timeSleepConnect = model.DefaultTimeSleepConnect
			retryTimesConnect = model.DefaultRetryTimesConnect
			time.Sleep(time.Duration(model.DefaultTimeSleep) * time.Second)

			break

		handle_error:
			totalError++
			if totalError <= 18 {
				if totalError == 3 {
					// TODO push log to Senjiu system here
					time.Sleep(time.Duration(model.DefaultTimeSleepConnect) * time.Second)

					break
				}
				time.Sleep(time.Duration(model.ErrorTimeSleepConnect) * time.Second)
			} else {
				// TODO push log to Senjiu system here
				totalError = 3
				timeSleepConnect = model.ErrorTimeSleepConnect
				retryTimesConnect = model.ErrorRetryTimesConnect
				time.Sleep(time.Duration(model.ErrorTimeSleepConnect) * time.Second)
			}

			break
		}

	}
}
