package event

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func (h Handler) Delivery(ctx context.Context, msg []byte) error {
	err := h.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return fmt.Errorf("error while delivering message: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	_, _, err = h.conn.ReadMessage()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return fmt.Errorf("timeout error: no response received within 2 seconds")
		}
		return fmt.Errorf("error while reading message: %w", err)
	}
	return nil
}
