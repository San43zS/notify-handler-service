package service

import (
	notification2 "Notify-handler-service/internal/service/notification"
	"Notify-handler-service/internal/storage"
)

type Service interface {
	Notification() notification2.Notification
}

type service struct {
	storage storage.Storage
}

func New(repos storage.Storage) Service {
	return &service{
		storage: repos,
	}
}

func (s *service) Notification() notification2.Notification {
	return s.storage.Notification()
}
