package producer

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	"context"
	"github.com/op/go-logging"
	amqp "github.com/rabbitmq/amqp091-go"
)

var log = logging.MustGetLogger("producer")

type Producer interface {
	Produ(ctx context.Context, arr []byte) error
}

type producer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Producer {
	return &producer{
		dial: dial,
	}
}

func (p producer) Produ(ctx context.Context, arr []byte) error {
	body := arr
	err := p.dial.PublishWithContext(ctx,
		config.ProducerExchangeName, // exchange
		config.ProducerQueueName,    // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		log.Criticalf("failed to publish a message: %v", err)
		return err
	}

	return nil
}
