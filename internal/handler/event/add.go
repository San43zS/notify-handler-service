package event

import (
	mssg "Notify-handler-service/internal/handler/model/msg"
	message "Notify-handler-service/internal/handler/model/msg/parser/rabbitParser"
	"Notify-handler-service/internal/model/notification"
	"Notify-handler-service/pkg/encoding"
	"context"
	"fmt"
)

func (h Handler) Add(ctx context.Context, msg []byte) error {
	m, err := message.New().Parse(msg)
	if err != nil {
		return err
	}

	nID := encoding.New().NotificationID(m.UserId)

	notification := notification.Notification{
		Id:        nID,
		UserId:    m.UserId,
		Status:    mssg.CurrentStatus,
		Data:      string(m.Content),
		CreatedAt: m.CreatedAt,
		ExpiredAt: m.ExpiredAt,
	}

	err = h.srv.NotificationRedis().Add(ctx, notification)
	if err != nil {
		return fmt.Errorf("error while adding notification to redis: %w", err)
	}

	return nil
}
