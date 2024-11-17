package notifyRedis

import (
	notify "Notify-handler-service/internal/model/notification"
	"Notify-handler-service/internal/service/api/notification"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"fmt"
	"strconv"
)

type Notify struct {
	storage redis.Store
}

func New(storage redis.Store) notification.NotifyRedis {
	return &Notify{
		storage: storage,
	}
}

func (n Notify) Add(ctx context.Context, notification notify.Notification) error {
	err := n.storage.Cache().Set(ctx, notification.Id, "", notification.TTL)
	if err != nil {
		return fmt.Errorf("failed to add notification to redis: %w", err)
	}

	return nil
}

func (n Notify) Delete(ctx context.Context, id int) error {
	err := n.storage.Cache().Delete(ctx, strconv.Itoa(id))
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}

	return nil
}
