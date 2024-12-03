package server

import (
	"Notify-handler-service/internal/broker"
	handler "Notify-handler-service/internal/handler"
	"Notify-handler-service/internal/server/launcher"
	redisPubSub "Notify-handler-service/internal/server/launcher/pubSub"
	"Notify-handler-service/internal/server/launcher/rabbit"
	"Notify-handler-service/internal/server/launcher/websocket"
	"Notify-handler-service/internal/service"
	"context"
	"github.com/op/go-logging"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var log = logging.MustGetLogger("server")

type server struct {
	servers []launcher.Server
}

func New(srv service.Service, pubSub *redis.PubSub, broker broker.Broker) (launcher.Server, error) {
	wsConn, _ := websocket.New()

	h := handler.New(srv, broker, wsConn)

	s := &server{
		servers: []launcher.Server{
			rabbit.New(broker.RabbitMQ, h.Event),
			redisPubSub.New(pubSub, h.Event),
		},
	}

	return s, nil
}

func (s server) Serve(ctx context.Context) error {
	ctx, stop := context.WithCancel(ctx)
	errCh := make(chan error)

	gr, grCtx := errgroup.WithContext(ctx)

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
			log.Criticalf("server stopped with error: %v", err)
		}
	}

	stop()

	log.Infof("app: shutting down the server...")
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
