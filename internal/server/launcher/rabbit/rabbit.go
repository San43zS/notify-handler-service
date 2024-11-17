package rabbit

import (
	"Notify-handler-service/internal/broker/rabbit"
	"Notify-handler-service/internal/server/launcher"
	"Notify-handler-service/pkg/msghandler"
	"context"
	"fmt"
	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"
	"sync"
)

var log = logging.MustGetLogger("rabbit")

type server struct {
	handler msghandler.MsgResolver
	broker  rabbit.Service
}

func New(broker rabbit.Service, handler msghandler.MsgResolver) launcher.Server {
	server := &server{
		handler: handler,
		broker:  broker,
	}

	return server
}

func (s server) Serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(1)

	gr, grCtx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		defer wg.Done()

		return s.serve(grCtx)
	})

	wg.Wait()

	return gr.Wait()
}

func (s server) serve(ctx context.Context) error {
	c := s.broker.Consumer()
	for {
		if err := ctx.Err(); err != nil {
			log.Criticalf("Rabbit listener stopped error: %v", err)
			return fmt.Errorf("rabbit listener stopped error: %v", err)
		}

		m, err := c.Consume(ctx)
		if err != nil {
			log.Infof("failed to consume message error: %v", err)
			continue
		}

		go func() {
			err := s.handler.ServeMSG(ctx, m)
			if err != nil {
				log.Criticalf("failed to handle message: %v", err)
				return
			}
		}()
	}
}
