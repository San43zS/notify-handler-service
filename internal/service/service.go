package service

import (
	"Notify-handler-service/internal/handler/notification"
	"Notify-handler-service/internal/storage"
)

type Service interface {
	Notification() notification.Notification
}

type service struct {
	storage storage.Storage
}

func New(repos storage.Storage) Service {
	return &service{
		storage: repos,
	}
}

func (s *service) Notification() notification.Notification {
	return s.storage.Notification()
}
