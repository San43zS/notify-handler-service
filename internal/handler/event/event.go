package event

import (
	"Notify-handler-service/internal/broker/rabbit"
	message "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/event"
	"Notify-handler-service/internal/service"
	"Notify-handler-service/pkg/msghandler"
	"encoding/json"
	"fmt"
)

type Handler struct {
	srv    service.Service
	router msghandler.MsgHandler
}

func New(srv service.Service, broker rabbit.Service) msghandler.MsgHandler {
	eventParseFn := func(msg []byte) (string, error) {
		var common message.Common
		if err := json.Unmarshal(msg, &common); err != nil {
			return "", fmt.Errorf("error while parsing msg: %w", err)
		}
		return common.Type, nil
	}

	handler := Handler{
		srv:    srv,
		router: msghandler.New(eventParseFn),
	}

	handler.initHandler()

	return handler.router
}

func (h Handler) initHandler() {
	h.router.Add(event.AddNotify, h.Add)
	h.router.Add(event.ChangeExpired, h.ChangeExpired)
}
