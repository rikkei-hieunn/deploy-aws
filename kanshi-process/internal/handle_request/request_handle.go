// Package handlerequest for handle request
package handlerequest

import (
	"context"
	"fmt"
	"kanshi-process/configs"
	"kanshi-process/model"
	"strconv"
	"strings"
)

type requestHandler struct {
	config *configs.Server
}

// NewRequestHandler constructor init new request handler
func NewRequestHandler(cfg *configs.Server) IRequestHandler {
	return &requestHandler{
		config: cfg,
	}
}

// StartSendRequest start send request to toiawase and furiwake
func (r *requestHandler) StartSendRequest(ctx context.Context) error {
	request, err := FormatRequest(&r.config.TickSystem)
	if err != nil {
		return err
	}

	firstToiawaseCheck := NewWorker(request, &r.config.TickSocket)
	secondToiawaseCheck := NewWorker(request, &r.config.TickSocket)
	thirdToiawaseCheck := NewWorker(request, &r.config.TickSocket)
	furiwakeCheck := NewWorker(request, &r.config.TickSocket)

	go firstToiawaseCheck.Process(ctx, r.config.TickSocket.ToiawaseSocket.Host, r.config.TickSocket.ToiawaseSocket.Ports[0])
	go secondToiawaseCheck.Process(ctx, r.config.TickSocket.ToiawaseSocket.Host, r.config.TickSocket.ToiawaseSocket.Ports[1])
	go thirdToiawaseCheck.Process(ctx, r.config.TickSocket.ToiawaseSocket.Host, r.config.TickSocket.ToiawaseSocket.Ports[2])
	go furiwakeCheck.Process(ctx, r.config.TickSocket.FuriwakeSocket.Host, r.config.TickSocket.FuriwakeSocket.Port)

	return nil
}

// FormatRequest format request, append space character
func FormatRequest(config *configs.TickSystem) (configs.Request, error) {
	requestKinou := config.RequestKinou
	if len(requestKinou) != model.FunctionTypeLength || (requestKinou != model.FirstFunctionType && requestKinou != model.SecondFunctionType) {
		return nil, fmt.Errorf("invalid kinou")
	}

	requestKanriID := config.RequestKanriID
	if len(requestKanriID) < model.KanriIDLength {
		requestKanriID += strings.Repeat(model.SpaceString, model.KanriIDLength-len(requestKanriID))
	} else if len(requestKanriID) > model.KanriIDLength {
		return nil, fmt.Errorf("invalid kanri id")
	}

	requestUserID := config.RequestUserID
	if len(requestUserID) < model.UserIDLength {
		requestUserID += strings.Repeat(model.SpaceString, model.UserIDLength-len(requestUserID))
	} else if len(requestUserID) > model.UserIDLength {
		return nil, fmt.Errorf("invalid user id")
	}

	requestSyubetu := config.RequestSyubetu
	if len(requestSyubetu) != model.SyubetuLength {
		return nil, fmt.Errorf("invalid syubetu")
	}

	requestQuoteCode := config.RequestQuoteCode
	if len(requestQuoteCode) < model.QuoteCodeLength {
		requestQuoteCode += strings.Repeat(model.SpaceString, model.QuoteCodeLength-len(requestQuoteCode))
	} else if len(requestQuoteCode) > model.QuoteCodeLength {
		return nil, fmt.Errorf("invalid quote code")
	}

	requestFromDate := config.RequestFromDate
	if len(requestFromDate) < model.FromDateLength {
		requestFromDate += strings.Repeat(model.SpaceString, model.FromDateLength-len(requestFromDate))
	} else if len(requestFromDate) > model.FromDateLength {
		return nil, fmt.Errorf("invalid from date")
	}

	requestToDate := config.RequestToDate
	if len(requestToDate) < model.ToDateLength {
		requestToDate += strings.Repeat(model.SpaceString, model.ToDateLength-len(requestToDate))
	} else if len(requestToDate) > model.ToDateLength {
		return nil, fmt.Errorf("invalid to date")
	}

	var requestFromTime, requestToTime, requestYobi string
	if requestKinou == model.SecondFunctionType {
		requestFromTime = config.RequestFromTime
		if len(requestFromTime) < model.FromTimeLength {
			requestFromTime = strings.Repeat(model.SpaceString, model.FromTimeLength-len(requestFromTime)) + requestFromTime
		} else if len(requestFromTime) > model.FromTimeLength {
			return nil, fmt.Errorf("invalid from time")
		}

		requestToTime = config.RequestToTime
		if len(requestToTime) < model.ToTimeLength {
			requestToTime = strings.Repeat(model.SpaceString, model.ToTimeLength-len(requestToTime)) + requestToTime
		} else if len(requestToTime) > model.ToTimeLength {
			return nil, fmt.Errorf("invalid to time")
		}

		requestYobi = config.RequestYobi
		if len(requestYobi) < model.YobiLength {
			requestYobi += strings.Repeat(model.SpaceString, model.YobiLength-len(requestYobi))
		} else if len(requestYobi) > model.YobiLength {
			return nil, fmt.Errorf("invalid yobi")
		}
	}

	requestFunasi := config.RequestFunasi
	if len(requestFunasi) < model.MinuteFunasiLength {
		requestFunasi = strings.Repeat(model.SpaceString, model.MinuteFunasiLength-len(requestFunasi)) + requestFunasi
	} else if len(requestFunasi) > model.MinuteFunasiLength {
		return nil, fmt.Errorf("invalid funasi")
	}

	requestKikan := config.RequestKikan
	if len(requestKikan) < model.KikanLength {
		requestKikan += strings.Repeat(model.SpaceString, model.KikanLength-len(requestKikan))
	} else if len(requestKikan) > model.KikanLength {
		return nil, fmt.Errorf("invalid kikan")
	}

	requestElementsLength := strconv.Itoa(len(config.RequestElements))
	if len(requestElementsLength) < model.ElementsLength {
		requestElementsLength = strings.Repeat(model.SpaceString, model.KikanLength-len(requestElementsLength)) + requestElementsLength
	}

	return []string{
		requestKinou,
		requestKanriID,
		requestUserID,
		requestSyubetu,
		requestQuoteCode,
		requestFromDate,
		requestToDate,
		requestFromTime,
		requestToTime,
		requestYobi,
		requestFunasi,
		requestKikan,
		requestElementsLength,
		config.RequestElements,
	}, nil
}
