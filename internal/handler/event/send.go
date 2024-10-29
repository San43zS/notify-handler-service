package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"context"
	"fmt"
)

func (h *handler) Send() error {
	msg, err := h.srv.Notification().Send()
	if err != nil {
		return fmt.Errorf("error while receiving message from channel: %w", err)
	}

	arr, err := msg2.New().Unparse(msg)

	err = h.respondConsumer.p.Produce(context.Background(), arr)
	if err != nil {
		return fmt.Errorf("error while sending message to third service channel: %w", err)
	}
	return nil
}
