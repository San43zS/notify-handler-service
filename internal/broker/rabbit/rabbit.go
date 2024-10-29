package rabbit

import (
	"Notify-handler-service/internal/broker/rabbit/config"
	"Notify-handler-service/internal/broker/rabbit/consumer"
	"Notify-handler-service/internal/broker/rabbit/producer"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Producer() producer.Producer
	Consumer() consumer.Consumer
}

type service struct {
	dial *amqp.Channel
}

func New() (Service, error) {
	cfg := config.NewConfig()

	conn, err := amqp.Dial(cfg.Driver + cfg.URL)
	if err != nil {
		conn.Close()
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	if err := ConfigureConsumer(ch); err != nil {
		ch.Close()
		conn.Close()
	}

	if err := ConfigureProducer(ch); err != nil {
		ch.Close()
		conn.Close()
	}

	srv := &service{dial: ch}

	return srv, nil
}

func (s service) Producer() producer.Producer {
	return producer.New(s.dial)
}

func (s service) Consumer() consumer.Consumer {
	return consumer.New(s.dial)
}
