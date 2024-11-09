package msghandler

import (
	"context"
	"errors"
)

type EventParser func(msg []byte) (string, error)
type HandlerFunc1 func(ctx context.Context, msg []byte) error
type HandlerFunc2 func(ctx context.Context) ([]byte, error)

type MsgResolver interface {
	ServeMSG(ctx context.Context, msg []byte) error
}

type MsgHandler interface {
	MsgResolver
	Add(event string, fn interface{})
}

type handler struct {
	eventParser EventParser
	handlers    map[string]interface{}
}

func New(parser EventParser) MsgHandler {
	return &handler{
		eventParser: parser,
		handlers:    make(map[string]interface{}),
	}
}

func (h *handler) ServeMSG(ctx context.Context, msg []byte) error {
	event, err := h.eventParser(msg)
	if err != nil {
		return err
	}

	fn, ok := h.handlers[event]
	if !ok {
		return err
	}
	switch fn := fn.(type) {
	case HandlerFunc1:
		return fn(ctx, msg)
	case HandlerFunc2:
		_, err := fn(ctx)
		return err
	default:
		return errors.New("unknown handler type")
	}
}

func (h *handler) Add(event string, fn interface{}) {
	h.handlers[event] = fn
}
