package redis

import (
	"Notify-handler-service/internal/storage/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func Configuration(ctx context.Context, conn *redis.Client) error {
	pubSub := conn.Subscribe(ctx, config.ChannelName)

	_, err := pubSub.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
