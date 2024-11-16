package service

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/service/api/notification"
	"Notify-handler-service/internal/service/notification/notifyPsql"
	notifyBroker "Notify-handler-service/internal/service/notification/notifyRabbit"
	notifyStore "Notify-handler-service/internal/service/notification/notifyRedis"
	"Notify-handler-service/internal/storage"
	"Notify-handler-service/internal/storage/db/psql"
	"Notify-handler-service/internal/storage/db/redis"
)

type Service interface {
	NotificationRedis() notification.NotifyRedis
	NotificationRabbit() notification.NotifyRabbit
	NotificationPsql() notification.NotifyPsql
}

type service struct {
	psql         psql.Store
	redis        redis.Store
	notifyRedis  notification.NotifyRedis
	notifyRabbit notification.NotifyRabbit
	notifyPsql   notification.NotifyPsql
}

func New(repos storage.Storage, broker broker.Broker) Service {
	return &service{
		redis:        repos,
		psql:         repos,
		notifyRedis:  notifyStore.New(repos),
		notifyRabbit: notifyBroker.New(broker),
		notifyPsql:   notifyPsql.New(repos),
	}
}

func (s service) NotificationRedis() notification.NotifyRedis {
	return s.notifyRedis
}

func (s service) NotificationRabbit() notification.NotifyRabbit {
	return s.notifyRabbit
}

func (s service) NotificationPsql() notification.NotifyPsql {
	return s.notifyPsql
}
