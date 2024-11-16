package repo

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, exp time.Duration) error
	SetIfExists(ctx context.Context, key string, value any, exp time.Duration) error
	Delete(ctx context.Context, key string) error
}
