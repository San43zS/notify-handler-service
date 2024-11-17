package pubSub

import (
	"Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/event"
	"Notify-handler-service/internal/handler/model/msg/parser/notifyParser"
)

func Configuration(m []byte) ([]byte, error) {
	data := msg.Expired{
		Type: event.ChangeExpired,
		Id:   string(m),
	}
	return notifyParser.New().Unparse(data)
}
