package notifyRedis

import (
	notification3 "Notify-handler-service/internal/model/notification"
	notification2 "Notify-handler-service/internal/service/api/notification"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"strconv"
)

type Notify struct {
	storage redis.Store
}

func New(storage redis.Store) notification2.NotifyRedis {
	return &Notify{
		storage: storage,
	}
}

func (n Notify) Add(ctx context.Context, notification notification3.Notification) error {
	err := n.storage.Cache().Set(ctx, strconv.Itoa(notification.UserId), notification.Data, notification.TTL)

	if err != nil {
		return err
	}
	return nil
}

func (n Notify) Delete(ctx context.Context, id int) error {
	err := n.storage.Cache().Delete(ctx, strconv.Itoa(id))
	if err != nil {
		return err
	}
	return nil
}
