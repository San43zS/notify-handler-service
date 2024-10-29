package rabbit

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConfigureConsumer(ch *amqp.Channel) error {

	err := ch.ExchangeDeclare(
		config.ConExchangeName, // name
		config.ExchangeType,    // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		config.ConQueueName, // name
		false,               // durable
		false,               // delete when unused
		true,                // exclusive
		false,               // no-wait
		nil,                 // arguments
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,                 // name
		config.ConQueueName,    // key
		config.ConExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func ConfigureProducer(ch *amqp.Channel) error {

	err := ch.ExchangeDeclare(
		config.ProExchangeName, // name
		config.ExchangeType,    // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		config.ProQueueName, // name
		false,               // durable
		false,               // delete when unused
		true,                // exclusive
		false,               // no-wait
		nil,                 // arguments
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,                 // name
		config.ProQueueName,    // key
		config.ProExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}
