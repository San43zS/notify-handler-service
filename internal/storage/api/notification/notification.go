package notification

import (
	message "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/model/notification"
	"context"
)

type Notification interface {
	Add(ctx context.Context, notification notification.Notification) error
	ChangeStatus(ctx context.Context, id string, status string) error

	GetCurrent(ctx context.Context, userID int) ([]message.Notify, error)

	GetOld(ctx context.Context, userID int) ([]message.Notify, error)
}
