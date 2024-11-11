package service

import (
	"Notify-handler-service/internal/service/api/notification"
	notification2 "Notify-handler-service/internal/service/notification"
	"Notify-handler-service/internal/storage/db/redis"
)

type Service interface {
	Notification() notification.Notification
}

type service struct {
	storage redis.Store
	notify  notification.Notification
}

func New(repos redis.Store) Service {
	return &service{
		storage: repos,
		notify:  notification2.New(repos),
	}
}

func (s service) Notification() notification.Notification {
	return s.notify
}
