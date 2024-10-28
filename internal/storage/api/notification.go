package api

import (
	notification3 "Notify-handler-service/internal/model/notification"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"strconv"
)

type Notification interface {
	Add(ctx context.Context, notification notification3.Notification) error
	Delete(ctx context.Context, id int) error
}

type notification struct {
	redis.Store
}

func (n notification) Add(ctx context.Context, notification notification3.Notification) error {
	err := n.Cache().Set(ctx, strconv.Itoa(notification.UserId), notification.Data, notification.TTL)
	if err != nil {
		return err
	}
	return nil
}

func (n notification) Delete(ctx context.Context, id int) error {
	err := n.Cache().Delete(ctx, strconv.Itoa(id))
	if err != nil {
		return err
	}
	return nil
}
