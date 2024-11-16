package psql

import (
	"Notify-handler-service/internal/storage/api/notification"
	n "Notify-handler-service/internal/storage/db/psql/notification"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store interface {
	Notification() notification.Notification
}

type store struct {
	notification notification.Notification
}

func New(config *Config) (Store, error) {
	db, err := sqlx.Connect(config.Driver, config.URL)
	if err != nil {
		return nil, err
	}

	return &store{
		notification: n.New(db),
	}, nil
}

func (s store) Notification() notification.Notification {
	return s.notification
}
