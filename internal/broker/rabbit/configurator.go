package rabbit

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Configure(ch *amqp.Channel) error {

	err := ch.ExchangeDeclare(
		config.ExchangeName, // name
		config.ExchangeType, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		config.QueueName, // name
		false,            // durable
		false,            // delete when unused
		true,             // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,              // name
		config.QueueName,    // key
		config.ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}
