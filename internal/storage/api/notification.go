package api

import "context"

type Notification interface {
	Add(ctx context.Context, notification Notification) error
	Delete(ctx context.Context, id int) error
}

type notification struct {
	Notification
}

func (n notification) Add(ctx context.Context, notification Notification) error {

	return nil
}

func (n notification) Delete(ctx context.Context, id int) error {
	return nil
}
