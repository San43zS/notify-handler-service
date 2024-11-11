package consumer

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	Consume(ctx context.Context) ([]byte, error)
}

type consumer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Consumer {
	return &consumer{
		dial: dial,
	}
}

func (c consumer) Consume(ctx context.Context) ([]byte, error) {

	msgs, err := c.dial.Consume(
		config.ConsumerQueueName, // queue
		"",                       // consumer
		false,                    // auto-ack
		false,                    // exclusive
		false,                    // no-local
		true,                     // no-wait
		nil,                      // args
	)
	if err != nil {

		return nil, err
	}

	for msg := range msgs {

		return msg.Body, nil
	}

	return nil, nil
}
