package api

import "context"

type Notification interface {
	Add(ctx context.Context, notification Notification) error
	Delete(ctx context.Context, id int) error
}
