package notifyRabbit

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/broker/rabbit/consumer"
	"Notify-handler-service/internal/broker/rabbit/producer"
	"context"
)

type RespCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(broker broker.Broker) RespCons {
	return RespCons{
		p: broker.RabbitMQ.Producer(),
		c: broker.RabbitMQ.Consumer(),
	}
}

func (r RespCons) Add(ctx context.Context) ([]byte, error) {
	consume, err := r.c.Consume(ctx)
	if err != nil {
		return nil, err
	}
	return consume, nil
}

func (r RespCons) AddExpired(ctx context.Context, msg []byte) error {
	err := r.p.Produce(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (r RespCons) GetCurrent(ctx context.Context) ([]byte, error) {
	return nil, nil
}

func (r RespCons) GetOld(ctx context.Context) ([]byte, error) {
	return nil, nil
}
