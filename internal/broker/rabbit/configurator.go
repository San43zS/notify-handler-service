package rabbit

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConfigureConsumer(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		config.ConsumerExchangeName, // name
		amqp.ExchangeDirect,         // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		config.ConsumerQueueName, // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,                      // name
		config.ConsumerRoutingKey,   // key
		config.ConsumerExchangeName, // exchange
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
		config.ProducerExchangeName, // name
		amqp.ExchangeDirect,         // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		config.ProducerQueueName, // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,                      // name
		config.ProducerRoutingKey,   // key
		config.ProducerExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
