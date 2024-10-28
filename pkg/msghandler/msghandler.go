package msghandler

import "context"

type HandlerFunc func(ctx context.Context, msg []byte) error

func (h HandlerFunc) ServeMSG(ctx context.Context, msg []byte) error {
	return nil
}

type MsgResolver interface {
	ServeMSG(ctx context.Context, msg []byte) error
}

type handler struct {
	handler HandlerFunc
}

func New(fn HandlerFunc) MsgResolver {
	return &handler{
		handler: fn,
	}
}

func (h *handler) ServeMSG(ctx context.Context, msg []byte) error {
	return h.handler(ctx, msg)
}
