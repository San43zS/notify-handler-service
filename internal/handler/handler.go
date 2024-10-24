package handler

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/service"
)

type Handler struct {
}

func New(srv service.Service, broker broker.Broker) *Handler {
	return &Handler{}
}
