package msghandler

import (
	"context"
)

type EventParser func(msg []byte) (string, error)
type (
	HandlerFunc1 func(ctx context.Context, msg []byte) error
)
type HandlerFunc2 func(ctx context.Context) ([]byte, error)

type MsgResolver interface {
	ServeMSG(ctx context.Context, msg []byte) error
}

type MsgHandler interface {
	MsgResolver
	Add(event string, fn HandlerFunc1)
}

type handler struct {
	eventParser EventParser
	handlers    map[string]HandlerFunc1
}

func New(parser EventParser) MsgHandler {
	return &handler{
		eventParser: parser,
		handlers:    make(map[string]HandlerFunc1),
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
	return fn(ctx, msg)
}

func (h *handler) Add(event string, fn HandlerFunc1) {
	h.handlers[event] = fn
}
