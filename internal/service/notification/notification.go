package notification

import (
	notification3 "Notify-handler-service/internal/model/notification"
	notification2 "Notify-handler-service/internal/service/api/notification"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"log"
	"strconv"
)

type Notify struct {
	storage redis.Store
}

func New(storage redis.Store) notification2.Notification {
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

func (n Notify) Send() (notification3.Notification, error) {
	conn := n.storage.PubSub()
	mCh := make(chan notification3.Notification, 1)
	errCh := make(chan error, 1)

	go func() {
		m, err := conn.Receive(context.Background())
		if err != nil {
			errCh <- err
			return
		}
		msg, ok := m.(notification3.Notification)
		if !ok {
			log.Fatal("unexpected type")
		}
		mCh <- msg
	}()

	select {
	case m := <-mCh:
		return m, nil
	case err := <-errCh:
		return notification3.Notification{}, err
	}
}
