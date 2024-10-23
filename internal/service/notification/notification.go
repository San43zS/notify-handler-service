package notification

import (
	"Notify-handler-service/internal/model/notification"
	"context"
)

type Notification interface {
	Add(ctx context.Context, notification notification.Notification) error
	Delete(ctx context.Context, id int) error
}
