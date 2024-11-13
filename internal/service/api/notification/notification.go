package notification

import (
	notification2 "Notify-handler-service/internal/model/notification"
	"context"
)

type NotifyRedis interface {
	Add(ctx context.Context, notification notification2.Notification) error
	Delete(ctx context.Context, id int) error
}

type NotifyRabbit interface {
	GetOld(ctx context.Context) ([]byte, error)
	GetCurrent(ctx context.Context) ([]byte, error)
	AddExpired(ctx context.Context, msg []byte) error
	Add(ctx context.Context) ([]byte, error)
}
