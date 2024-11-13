package service

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/service/api/notification"
	notifyBroker "Notify-handler-service/internal/service/notification/notifyRabbit"
	notifyStore "Notify-handler-service/internal/service/notification/notifyRedis"
	"Notify-handler-service/internal/storage/db/redis"
)

type Service interface {
	NotificationRedis() notification.NotifyRedis
	NotificationRabbit() notification.NotifyRabbit
}

type service struct {
	storage      redis.Store
	notifyRedis  notification.NotifyRedis
	notifyRabbit notification.NotifyRabbit
}

func New(repos redis.Store, broker broker.Broker) Service {
	return &service{
		storage:      repos,
		notifyRedis:  notifyStore.New(repos),
		notifyRabbit: notifyBroker.New(broker),
	}
}

func (s service) NotificationRedis() notification.NotifyRedis {
	return s.notifyRedis
}

func (s service) NotificationRabbit() notification.NotifyRabbit {
	return s.notifyRabbit
}
