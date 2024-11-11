package pubSub

import (
	"Notify-handler-service/internal/server/launcher"
	"Notify-handler-service/pkg/msghandler"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
)

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

// ???????????????????????????????????????????/
func (s server) serve(ctx context.Context) error {
	conn := s.pubSub
	for {
		if err := ctx.Err(); err != nil {
			err := fmt.Errorf("PubSub listener stopped error: %v", err)
			return err
		}

		m, err := conn.Receive(context.Background())
		if err != nil {
			log.Println("failed to receive message from pubSub:", err)
			continue
		}

		if msg, ok := m.(*redis.Message); ok {
			go func() {
				err := s.handler.ServeMSG(ctx, []byte(msg.Payload))
				if err != nil {
					fmt.Errorf("failed to handle message: %v", err)
					return
				}
			}()
		}
	}
}
