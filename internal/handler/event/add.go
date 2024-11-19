package event

import (
	mssg "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/parser/rabbitParser/messageParser"
	"Notify-handler-service/internal/model/notification"
	"Notify-handler-service/pkg/encoding"
	"context"
	"fmt"
)

func (h Handler) Add(ctx context.Context, msg []byte) error {
	m, err := messageParser.New().Parse(msg)
	if err != nil {
		return err
	}

	nID := encoding.New().NotificationID(m.Data.UserId)

	notification := notification.Notification{
		Id:        nID,
		UserId:    m.Data.UserId,
		Status:    mssg.CurrentStatus,
		Data:      string(m.Data.Content),
		CreatedAt: m.Data.CreatedAt,
		ExpiredAt: m.Data.ExpiredAt,
	}

	err = h.srv.NotificationPsql().Add(ctx, notification)
	if err != nil {
		return fmt.Errorf("error while adding notification to psql: %w", err)
	}

	err = h.srv.NotificationRedis().Add(ctx, notification)
	if err != nil {
		return fmt.Errorf("error while adding notification to redis: %w", err)
	}

	return nil
}
