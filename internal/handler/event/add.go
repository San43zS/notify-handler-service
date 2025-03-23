package event

import (
	mssg "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/parser/rabbitParser/messageParser"
	"Notify-handler-service/internal/model/notification"
	"Notify-handler-service/pkg/encoding"
	"context"
	"fmt"
	"time"
)

func (h Handler) Add(ctx context.Context, msg []byte) error {

	m, err := messageParser.New().Parse(msg)
	if err != nil {
		return err
	}

	nID := encoding.New().NotificationID(1)

	timeNOW := time.Now()
	timeEXP := timeNOW.Add(time.Duration(m.TTL) * time.Second)

	n := notification.Notification{
		Id:        nID,
		UserId:    43,
		Status:    mssg.CurrentStatus,
		Data:      string(m.Data),
		TTL:       time.Duration(m.TTL) * time.Second,
		CreatedAt: timeNOW,
		ExpiredAt: timeEXP,
	}

	err = h.srv.NotificationPsql().Add(ctx, n)
	if err != nil {
		return fmt.Errorf("error while adding notification to psql: %w", err)
	}

	err = h.srv.NotificationRedis().Add(ctx, n)
	if err != nil {
		return fmt.Errorf("error while adding notification to redis: %w", err)
	}

	return nil
}
