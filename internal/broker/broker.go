package broker

import (
	"Notify-handler-service/internal/broker/rabbit"
	"context"
	"fmt"
)

type Broker struct {
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

// TODO: add config for array of consumers!!!!!!!!
func (b Broker) Start(ctx context.Context) error {

	return nil
}
