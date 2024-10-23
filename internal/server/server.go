package server

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/service"
)

type server struct {
	broker broker.Broker
}

func New(srv service.Service) (broker.Broker, error) {
	brk, err := broker.New()
	if err != nil {
		return broker.Broker{}, err
	}

	return brk, nil
}

//TODO func Run() error
