package api

import (
	notification3 "Notify-handler-service/internal/model/notification"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"log"
	"strconv"
)

type Notification interface {
	Add(ctx context.Context, notification notification3.Notification) error
	Delete(ctx context.Context, id int) error
	Send() (notification3.Notification, error)
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

func (n notification) Send() (notification3.Notification, error) {
	conn := n.PubSub()
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
