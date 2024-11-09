package event

import (
	"context"
	"fmt"
)

func (h Handler) AddNotify(ctx context.Context, msg []byte) error {

	recMSG, err := InitNotify(msg)
	if err != nil {
		return fmt.Errorf("error while initializing notification: %w", err)
	}

	err = h.srv.Notification().Add(ctx, recMSG)
	if err != nil {
		return fmt.Errorf("error while adding notification to redis: %w", err)
	}

	return nil
}
