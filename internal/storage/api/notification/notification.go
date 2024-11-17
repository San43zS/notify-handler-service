package notification

import (
	"Notify-handler-service/internal/model/notification"
	"context"
)

type Notification interface {
	Add(ctx context.Context, notification notification.Notification) error
	ChangeStatus(ctx context.Context, id string, status string) error
}
