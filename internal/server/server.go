package server

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/service"
	"context"
	"fmt"
)

type Server struct {
	broker broker.Broker
}

func New(srv service.Service) (*Server, error) {
	brk, err := broker.New()
	if err != nil {
		return &Server{
			broker: brk,
		}, err
	}

	return &Server{
		broker: brk,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	err := s.broker.Start(ctx)
	if err != nil {
		return fmt.Errorf("failed to start broker: %w", err)
	}
	return nil
}
