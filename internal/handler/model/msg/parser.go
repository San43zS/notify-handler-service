package msg

import (
	notification "Notify-handler-service/internal/model/notification"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse() (notification.Notification, error)
}

type parser struct {
	data []byte
}

func New(data []byte) Parser {
	return &parser{
		data: data,
	}
}

func (p parser) Parse() (notification.Notification, error) {
	var msg MSG
	if err := json.Unmarshal(p.data, &msg); err != nil {
		return notification.Notification{}, fmt.Errorf("error while parsing msg: %w", err)
	}
	return notification.Notification{UserId: msg.UserId, Data: msg.Data, TTL: msg.TTL, CreatedAt: msg.CreatedAt}, nil
}
