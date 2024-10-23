package redis

import (
	"Notify-handler-service/internal/model/notification"
	notification2 "Notify-handler-service/internal/service/notification"
	"Notify-handler-service/internal/storage/config"
	cashRep "Notify-handler-service/internal/storage/db/redis/cache"
	"Notify-handler-service/internal/storage/repo"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Store interface {
	Cache() repo.Cache
	Close() error
	Notification() notification2.Notification
}

type store struct {
	cache  repo.Cache
	db     *redis.Client
	notify notification.Notification
}

func configure(db *redis.Client) Store {
	return store{
		cache: cashRep.New(db),
		db:    db,
	}
}

func New(config config.Config) (Store, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.URL,
		Password: config.Password,
		Username: config.Username,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	err = Configuration(context.Background(), client)
	if err != nil {
		log.Fatal("Error while subscribing to redis channel: ", err)
		return nil, err
	}

	return configure(client), nil
}

func (s store) Close() error {
	return s.db.Close()
}

func (s store) Cache() repo.Cache {
	return s.cache
}

func (s store) Notification() notification2.Notification {
	return s.Notification()
}
