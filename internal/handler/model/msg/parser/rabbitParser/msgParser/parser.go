package msgParser

import (
	"Notify-handler-service/internal/handler/model/msg"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) (msg.MSG, error)
	Unparse(msg.MSG) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) (msg.MSG, error) {
	var message msg.MSG
	test := string(m)
	if err := json.Unmarshal([]byte(test), &message); err != nil {
		return msg.MSG{}, fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	return message, nil
}

func (p parser) Unparse(m msg.MSG) ([]byte, error) {
	arr, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while parsing(marshal) msg: %w", err)
	}
	return arr, nil
}
