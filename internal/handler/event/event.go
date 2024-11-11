package event

import (
	"Notify-handler-service/internal/broker/rabbit"
	"Notify-handler-service/internal/broker/rabbit/consumer"
	"Notify-handler-service/internal/broker/rabbit/producer"
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/event"
	"Notify-handler-service/internal/service"
	"Notify-handler-service/pkg/msghandler"
	"fmt"
)

type Handler struct {
	srv    service.Service
	router msghandler.MsgHandler

	respondConsumer respCons
}

type respCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(srv service.Service, broker rabbit.Service) msghandler.MsgHandler {
	eventParseFn := func(msg []byte) (string, error) {
		m, err := msg2.New().Parse(msg)
		if err != nil {
			return "", fmt.Errorf("error while parsing msg: %w", err)
		}
		return m.Type, nil
	}

	handler := Handler{
		srv:    srv,
		router: msghandler.New(eventParseFn),
	}

	handler.initHandler()

	return handler.router
}

func (h Handler) initHandler() {
	//h.router.Add(event.SendNotify, h.SendNotify)
	h.router.Add(event.AddNotify, h.AddNotify)
}
