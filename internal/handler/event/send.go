package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"context"
	"fmt"
)

func (h Handler) SendNotify() error {
	msg, err := h.srv.Notification().Send()
	if err != nil {
		return fmt.Errorf("error while receiving message from channel: %w", err)
	}
	result := msg2.Message{
		UserId:    msg.UserId,
		CreatedAt: msg.CreatedAt,
		Data: msg2.Data{
			Data: []byte(msg.Data),
		},
	}

	arr, err := msg2.New().Unparse(result)

	err = h.respondConsumer.p.Produce(context.Background(), arr)
	if err != nil {
		return fmt.Errorf("error while sending message to third service channel: %w", err)
	}
	return nil
}
