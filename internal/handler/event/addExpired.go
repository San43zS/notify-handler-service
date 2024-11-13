package event

import (
	"Notify-handler-service/internal/handler/model/msg/parser/notifyParser"
	"Notify-handler-service/internal/handler/model/msg/parser/rabbitParser"
	notification3 "Notify-handler-service/internal/model/notification"
	"context"
	"time"
)

func (h Handler) AddExpired(ctx context.Context, msg []byte) error {
	m, err := rabbitParser.New().Parse(msg)
	if err != nil {
		return err
	}

	notification := notification3.Notify{
		Number:    0,
		Data:      string(m.Content.Data),
		CreatedAt: time.Now(),
	}

	data, err := notifyParser.New().Unparse(notification)
	if err != nil {
		return err
	}

	err = h.srv.NotificationRabbit().AddExpired(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
