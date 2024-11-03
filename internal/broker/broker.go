package broker

import (
	"Notify-handler-service/internal/broker/rabbit"
	"Notify-handler-service/pkg/msghandler"
	"fmt"
)

type Broker struct {
	handler  msghandler.MsgResolver
	RabbitMQ rabbit.Service
}

func New() (Broker, error) {
	rabbitMQ, err := rabbit.New()
	if err != nil {

		return Broker{}, fmt.Errorf("failed to create RabbitMQ broker: %w", err)
	}

	broker := Broker{
		RabbitMQ: rabbitMQ,
	}

	return broker, nil
}
