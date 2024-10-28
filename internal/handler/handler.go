package handler

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/handler/event"
	"Notify-handler-service/internal/service"
	"Notify-handler-service/pkg/msghandler"
)

type Handler struct {
	Event msghandler.MsgResolver
}

func New(srv service.Service, brk broker.Broker) *Handler {
	return &Handler{
		Event: event.New(srv, brk.RabbitMQ),
	}
}
