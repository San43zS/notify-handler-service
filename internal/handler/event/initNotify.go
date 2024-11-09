package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/event"
	"Notify-handler-service/internal/model/notification"
	"fmt"
	"time"
)

func InitNotify(msg []byte) (notification.Notification, error) {
	m, err := msg2.New().Parse(msg)
	if err != nil {
		return notification.Notification{}, fmt.Errorf("error while parsing(initNotify) msg: %w", err)
	}
	result := notification.Notification{
		UserId:    event.User_ID,
		Data:      string(m.Content.Data),
		TTL:       event.TTL,
		CreatedAt: time.Now(),
	}
	return result, nil
}
