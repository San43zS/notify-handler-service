package server

import (
	"Notify-handler-service/internal/broker"
	handler "Notify-handler-service/internal/handler"
	"Notify-handler-service/internal/server/launcher"
	"Notify-handler-service/internal/server/launcher/rabbit"
	"Notify-handler-service/internal/service"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type server struct {
	servers []launcher.Server
}

func New(srv service.Service) (launcher.Server, error) {
	brk, err := broker.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create broker: %w", err)
	}

	h := handler.New(srv, brk)

	s := &server{
		servers: []launcher.Server{
			rabbit.New(brk.RabbitMQ, h.Event),
		},
	}

	return s, nil
}

func (s server) Serve(ctx context.Context) error {
	ctx, stop := context.WithCancel(ctx)
	errCh := make(chan error)

	gr, grCtx := errgroup.WithContext(ctx)

	// start server
	gr.Go(func() error {
		return s.serve(grCtx)
	})

	go func() {
		defer close(errCh)
		errCh <- gr.Wait()
	}()

	var err error

	select {
	case <-getExitSignal():

	case err = <-errCh:
		if err != nil {
			fmt.Errorf("app error: %v", err)
		}
	}

	stop()

	fmt.Println("app: shutting down the server...")
	<-errCh

	return err
}

func (s *server) serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, server := range s.servers {
		server := server

		gr.Go(func() error {
			defer wg.Done()

			return server.Serve(grCtx)
		})
	}

	wg.Wait()

	return gr.Wait()
}

func getExitSignal() <-chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return quit
}
