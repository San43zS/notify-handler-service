package event

import (
	"Notify-handler-service/internal/broker/rabbit"
	"Notify-handler-service/internal/broker/rabbit/consumer"
	"Notify-handler-service/internal/broker/rabbit/producer"
	"Notify-handler-service/internal/service"
	"Notify-handler-service/pkg/msghandler"
	"context"
)

type handler struct {
	srv    service.Service
	router msghandler.MsgResolver

	respondConsumer respCons
}

type respCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(srv service.Service, broker rabbit.Service) msghandler.MsgResolver {
	eventParseFn := func(ctx context.Context, msg []byte) error {
		//TODO: parse event
		return nil
	}

	handler := &handler{
		srv:    srv,
		router: msghandler.New(eventParseFn),
	}
	return handler.router
}
