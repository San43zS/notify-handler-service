package cache

import (
	repo2 "Notify-handler-service/internal/storage/db/repo"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
	ErrNotEqual = errors.New("not equal")

	ErrTransactionFailed = errors.New("transaction failed")
)

const (
	requestTimeout = 15 * time.Second
)

type repo struct {
	db *redis.Client
}

func New(db *redis.Client) repo2.Cache {
	return repo{
		db: db,
	}
}

func (r repo) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	val, err := r.db.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrNotFound

		}

		return "", err
	}

	return val, nil
}

func (r repo) Set(ctx context.Context, key string, value any, exp time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	return r.db.Set(ctx, key, value, exp).Err()
}

func (r repo) SetIfExists(ctx context.Context, key string, value any, exp time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	return r.db.SetXX(ctx, key, value, exp).Err()
}

func (r repo) Delete(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	if err := r.db.Del(ctx, key).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrNotFound
		}

		return err
	}

	return nil
}
