package notifyParser

import (
	"Notify-handler-service/internal/model/notification"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) (notification.Notify, error)
	Unparse(notification.Notify) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) (notification.Notify, error) {
	var message notification.Notify
	test := string(m)
	if err := json.Unmarshal([]byte(test), &message); err != nil {
		return notification.Notify{}, fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	return message, nil
}

func (p parser) Unparse(m notification.Notify) ([]byte, error) {
	arr, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while parsing(marshal) msg: %w", err)
	}
	return arr, nil
}
