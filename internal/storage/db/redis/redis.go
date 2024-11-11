package redis

import (
	"Notify-handler-service/internal/storage/config"
	cashRep "Notify-handler-service/internal/storage/db/redis/cache"
	"Notify-handler-service/internal/storage/repo"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Store interface {
	Cache() repo.Cache
	PubSub() *redis.PubSub
	Close() error
}

type store struct {
	cache repo.Cache
	db    *redis.Client
	conn  *redis.PubSub
}

func configure(db *redis.Client, c *redis.PubSub) Store {
	return store{
		cache: cashRep.New(db),
		db:    db,
		conn:  c,
	}
}

func New(config config.Config) (Store, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.URL,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	conn, err := Configuration(context.Background(), client)
	if err != nil {
		log.Fatal("Error while subscribing to redis channel: ", err)
		return nil, err
	}

	return configure(client, conn), nil
}

func (s store) Close() error {
	return s.db.Close()
}

func (s store) PubSub() *redis.PubSub {
	return s.conn
}

func (s store) Cache() repo.Cache {
	return s.cache
}
