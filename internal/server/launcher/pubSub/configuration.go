package pubSub

import (
	"Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/event"
	"Notify-handler-service/internal/handler/model/msg/parser/rabbitParser"
)

func Configuration(mag []byte) ([]byte, error) {
	m := msg.MSG{
		Type: event.AddExpired,
		Content: msg.Data{
			Data: mag,
		},
	}
	return rabbitParser.New().Unparse(m)
}
