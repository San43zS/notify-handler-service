package event

import (
	"Notify-handler-service/internal/broker/rabbit"
	message "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/event"
	"Notify-handler-service/internal/service"
	"Notify-handler-service/pkg/msghandler"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Handler struct {
	srv    service.Service
	router msghandler.MsgHandler
	conn   *websocket.Conn
}

func New(srv service.Service, broker rabbit.Service, ws *websocket.Conn) msghandler.MsgHandler {
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
		conn:   ws,
	}

	handler.initHandler()

	return handler.router
}

func (h Handler) initHandler() {
	h.router.Add(event.AddNotify, h.Add)
	h.router.Add(event.ChangeExpired, h.ChangeExpired)
}
