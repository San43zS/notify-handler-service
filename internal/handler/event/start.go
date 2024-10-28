package event

import (
	"context"
	"fmt"
)

func (h *handler) Start() error {
	consume, err := h.respondConsumer.c.Consume(context.Background())
	if err != nil {
		return fmt.Errorf("error while consuming: %w", err)
	}

	return h.Receive(context.Background(), consume)
}
