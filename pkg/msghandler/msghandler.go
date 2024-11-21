package msghandler

import (
	"context"
)

type EventParser func(msg []byte) (string, error)
type HandlerFunc func(ctx context.Context, msg []byte) error

type MsgResolver interface {
	ServeMSG(ctx context.Context, msg []byte) error
}

type MsgHandler interface {
	MsgResolver
	Add(event string, fn HandlerFunc)
}

type handler struct {
	eventParser EventParser
	handlers    map[string]HandlerFunc
}

func New(parser EventParser) MsgHandler {
	return &handler{
		eventParser: parser,
		handlers:    make(map[string]HandlerFunc),
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

func (h *handler) Add(event string, fn HandlerFunc) {
	h.handlers[event] = fn
}
