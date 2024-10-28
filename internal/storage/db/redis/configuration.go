package redis

import (
	"Notify-handler-service/internal/storage/config"
	"context"
	"github.com/redis/go-redis/v9"
)

func Configuration(ctx context.Context, conn *redis.Client) (*redis.PubSub, error) {
	pubSub := conn.Subscribe(ctx, config.ChannelName)
	return pubSub, nil
}
