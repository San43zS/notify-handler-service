package msg

import (
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) (MSG, error)
	Unparse(Message) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) (MSG, error) {
	var msg MSG
	if err := json.Unmarshal(m, &msg); err != nil {
		return MSG{}, fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	return msg, nil
}

func (p parser) Unparse(m Message) ([]byte, error) {
	arr, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while parsing(marshal) msg: %w", err)
	}
	return arr, nil
}
