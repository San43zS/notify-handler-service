package pubSub

import (
	"Notify-handler-service/internal/server/launcher"
	"Notify-handler-service/pkg/msghandler"
	"context"
	"fmt"
	"github.com/op/go-logging"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"sync"
)

var log = logging.MustGetLogger("pubsub")

type server struct {
	handler msghandler.MsgResolver
	pubSub  *redis.PubSub
}

func New(pubSub *redis.PubSub, handler msghandler.MsgResolver) launcher.Server {
	server := &server{
		handler: handler,
		pubSub:  pubSub,
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
	conn := s.pubSub
	for {
		if err := ctx.Err(); err != nil {
			log.Criticalf("PubSub listener stopped error: %v", err)
			return fmt.Errorf("PubSub listener stopped error: %v", err)
		}

		m, err := conn.Receive(context.Background())
		if err != nil {
			log.Criticalf("failed to receive message from pubSub: %v", err)
			continue
		}

		if msg, ok := m.(*redis.Message); ok {
			go func() {
				test := msg.Payload
				message, err := Configuration([]byte(test))
				if err != nil {
					log.Criticalf("failed to parse message: %v", err)
					return
				}

				err = s.handler.ServeMSG(ctx, message)
				if err != nil {
					log.Criticalf("failed to handle message: %v", err)
					return
				}
			}()
		}
	}
}
