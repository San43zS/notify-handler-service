package notification

import (
	notification3 "Notify-handler-service/internal/model/notification"
	"Notify-handler-service/internal/storage/api"
	"context"
)

type service struct {
	storage api.Notification
}

func New(repos api.Notification) api.Notification {
	return &service{
		storage: repos,
	}
}

func (s *service) Add(ctx context.Context, notification notification3.Notification) error {
	return s.storage.Add(ctx, notification)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.storage.Delete(ctx, id)
}

func (s *service) Send() (notification3.Notification, error) {
	return s.storage.Send()
}
