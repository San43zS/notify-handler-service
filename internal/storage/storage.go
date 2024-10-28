package storage

import (
	"Notify-handler-service/internal/handler/notification"
)

type Storage interface {
	Notification() notification.Notification
}
