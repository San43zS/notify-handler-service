package notification

import (
	"Notify-handler-service/internal/model/notification"
	"context"
)

type NotifyRedis interface {
	Add(ctx context.Context, notification notification.Notification) error
	Delete(ctx context.Context, id int) error
}

type NotifyRabbit interface {
	GetOld(ctx context.Context) ([]byte, error)
	GetCurrent(ctx context.Context) ([]byte, error)
	AddExpired(ctx context.Context, msg []byte) error
	Add(ctx context.Context) ([]byte, error)
}

type NotifyPsql interface {
	Add(ctx context.Context, notification notification.Notification) error
	ChangeStatus(ctx context.Context, id string, status string) error
}
