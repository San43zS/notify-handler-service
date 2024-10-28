package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"context"
	"fmt"
)

func (h *handler) Send() error {
	msg, err := h.srv.Notification().Receive()
	if err != nil {
		return fmt.Errorf("error while receiving message from channel: %w", err)
	}

	arr, err := msg2.New().Unparse(msg)
	h.respondConsumer.p.Produce(context.Background(), arr)
	return nil
}
