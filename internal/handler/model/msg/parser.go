package msg

import (
	notification "Notify-handler-service/internal/model/notification"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) (notification.Notification, error)
	Unparse(notification.Notification) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) (notification.Notification, error) {
	var msg MSG
	if err := json.Unmarshal(m, &msg); err != nil {
		return notification.Notification{}, fmt.Errorf("error while parsing msg: %w", err)
	}
	return notification.Notification{UserId: msg.UserId, Data: msg.Data, TTL: msg.TTL, CreatedAt: msg.CreatedAt}, nil
}

func (p parser) Unparse(notification notification.Notification) ([]byte, error) {
	var arr []byte
	arr, err := json.Marshal(notification)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling msg: %w", err)
	}
	return arr, nil
}
