package notifyPsql

import (
	message "Notify-handler-service/internal/handler/model/msg"
	notify "Notify-handler-service/internal/model/notification"
	"Notify-handler-service/internal/storage/api/notification"
	"Notify-handler-service/internal/storage/db/psql"
	"context"
)

type Notify struct {
	storage psql.Store
}

func New(storage psql.Store) notification.Notification {
	return &Notify{
		storage: storage,
	}
}

func (n Notify) Add(ctx context.Context, notification notify.Notification) error {
	return n.storage.Notification().Add(ctx, notification)
}

func (n Notify) ChangeStatus(ctx context.Context, id string, status string) error {
	return n.storage.Notification().ChangeStatus(ctx, id, status)
}

func (n Notify) GetCurrent(ctx context.Context, userID int) ([]message.Notify, error) {
	return n.storage.Notification().GetCurrent(ctx, userID)
}

func (n Notify) GetOld(ctx context.Context, userID int) ([]message.Notify, error) {
	return n.storage.Notification().GetOld(ctx, userID)
}
