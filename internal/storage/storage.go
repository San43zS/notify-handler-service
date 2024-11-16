package storage

import (
	"Notify-handler-service/internal/storage/api/cache"
	"Notify-handler-service/internal/storage/api/notification"
	"Notify-handler-service/internal/storage/db/psql"
	"Notify-handler-service/internal/storage/db/redis"
	r "github.com/redis/go-redis/v9"
)

type Storage interface {
	Notification() notification.Notification
	Cache() cache.Cache
	PubSub() *r.PubSub
	Close() error
}

type storage struct {
	psql  psql.Store
	redis redis.Store
}

func New() (Storage, error) {
	var err error
	var storage storage

	storage.psql, err = psql.New(psql.NewConfig())
	if err != nil {
		return nil, err
	}

	storage.redis, err = redis.New(redis.NewConfig())
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s storage) Notification() notification.Notification {
	return s.psql.Notification()
}

func (s storage) Cache() cache.Cache {
	return s.redis.Cache()
}

func (s storage) PubSub() *r.PubSub {
	return s.redis.PubSub()
}

func (s storage) Close() error {
	return s.redis.Close()
}
