package notifyRabbit

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/broker/rabbit/consumer"
	"Notify-handler-service/internal/broker/rabbit/producer"
	"context"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("NotificationRabbit")

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
	//err := r.p.Produce(ctx, nil)
	//if err != nil {
	//	return nil, err
	//}
	//return consume, nil
	return nil, nil
}

func (r RespCons) AddExpired(ctx context.Context, msg []byte) error {

	err := r.p.Produ(ctx, msg)
	if err != nil {
		log.Criticalf("failed to publish a message: %v", err)
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
