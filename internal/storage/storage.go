package storage

import (
	notification2 "Notify-handler-service/internal/service/notification"
)

type Storage interface {
	Notification() notification2.Notification
}
