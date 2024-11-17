package notification

import (
	notify "Notify-handler-service/internal/model/notification"
	"Notify-handler-service/internal/storage/api/notification"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) notification.Notification {
	return repository{
		db: db,
	}
}

func (r repository) Add(ctx context.Context, notification notify.Notification) error {
	query := `INSERT INTO notify (id, user_id, status, notification, created_at, expired_at) VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, notification.Id, notification.UserId, notification.Status,
		notification.Data, notification.CreatedAt, notification.ExpiredAt)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}

func (r repository) ChangeStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE notify SET status = $1 WHERE id = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, status, id)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}
