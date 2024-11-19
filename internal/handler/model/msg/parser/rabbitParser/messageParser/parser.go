package messageParser

import (
	"Notify-handler-service/internal/handler/model/msg"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) (msg.Message, error)
	Unparse(msg.Message) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) (msg.Message, error) {
	var message msg.Message
	test := string(m)
	if err := json.Unmarshal([]byte(test), &message); err != nil {
		return msg.Message{}, fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	return message, nil
}

func (p parser) Unparse(m msg.Message) ([]byte, error) {
	arr, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while parsing(marshal) msg: %w", err)
	}
	return arr, nil
}
