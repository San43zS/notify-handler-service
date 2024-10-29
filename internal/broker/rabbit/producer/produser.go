package producer

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	Produce(ctx context.Context, arr []byte) error
}

type producer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Producer {
	return &producer{
		dial: dial,
	}
}

func (p producer) Produce(ctx context.Context, arr []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.ContextTimeOut)
	defer cancel()

	body := arr
	err := p.dial.PublishWithContext(ctx,
		config.ProExchangeName, // exchange
		config.ProQueueName,    // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {

		return err
	}

	return nil
}
