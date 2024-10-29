package notification

import (
	notification2 "Notify-handler-service/internal/model/notification"
	"context"
)

type Notification interface {
	Add(ctx context.Context, notification notification2.Notification) error
	Delete(ctx context.Context, id int) error
	Send() (notification2.Notification, error)
}

type notification struct {
	//Notification
}
