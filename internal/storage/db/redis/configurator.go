package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func Configuration(ctx context.Context, conn *redis.Client) (*redis.PubSub, error) {
	pubSub := conn.Subscribe(ctx, ChannelName)
	return pubSub, nil
}
