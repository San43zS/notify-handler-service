package notification

import "time"

type Notification struct {
	Id        int
	Info      string
	TTL       time.Duration
	CreatedAt int
}
