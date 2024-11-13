package notifyRabbit

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/broker/rabbit/consumer"
	"Notify-handler-service/internal/broker/rabbit/producer"
	"context"
)

type respCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(broker broker.Broker) respCons {
	return respCons{
		p: broker.RabbitMQ.Producer(),
		c: broker.RabbitMQ.Consumer(),
	}
}

func (r respCons) Add(ctx context.Context) ([]byte, error) {
	consume, err := r.c.Consume(ctx)
	if err != nil {
		return nil, err
	}
	return consume, nil
}

func (r respCons) AddExpired(ctx context.Context, msg []byte) error {
	err := r.p.Produce(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (r respCons) GetCurrent(ctx context.Context) ([]byte, error) {
	return nil, nil
}

func (r respCons) GetOld(ctx context.Context) ([]byte, error) {
	return nil, nil
}
