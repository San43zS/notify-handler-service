package notification

import (
	message "Notify-handler-service/internal/handler/model/msg"
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

const (
	CurrentStatus = "not_viewed"
	OldStatus     = "viewed"
)

func (r repository) GetCurrent(ctx context.Context, userID int) ([]message.Notify, error) {
	query := `SELECT user_id, notification, created_at, expired_at FROM notify WHERE user_id = $1 AND status = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, userID, CurrentStatus)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []message.Notify
	for rows.Next() {
		var existing message.Notify
		err := rows.Scan(
			&existing.UserId,
			&existing.Content,
			&existing.CreatedAt,
			&existing.ExpiredAt,
		)
		if err != nil {
			return []message.Notify{}, fmt.Errorf("failed to scan row: %w", err)
		}
		notifications = append(notifications, existing)
	}

	return notifications, nil
}

func (r repository) GetOld(ctx context.Context, userID int) ([]message.Notify, error) {
	query := `SELECT user_id, notification, created_at FROM notify WHERE user_id = $1 AND status = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, userID, OldStatus)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []message.Notify
	for rows.Next() {
		var existing message.Notify
		err := rows.Scan(
			&existing.UserId,
			&existing.Content,
			&existing.CreatedAt,
		)
		if err != nil {
			return []message.Notify{}, fmt.Errorf("failed to scan row: %w", err)
		}
		notifications = append(notifications, existing)
	}

	return notifications, nil
}
