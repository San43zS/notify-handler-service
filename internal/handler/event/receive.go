package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"context"
	"fmt"
)

func (h *handler) Receive(ctx context.Context, m []byte) error {

	recMSG, err := msg2.New().Parse(m)
	if err != nil {
		return fmt.Errorf("error while parsing msg: %w", err)
	}

	err = h.srv.Notification().Add(ctx, recMSG)
	if err != nil {
		return fmt.Errorf("error while adding notification: %w", err)
	}

	return nil
}
