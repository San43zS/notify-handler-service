package notification

import "time"

type Notification struct {
	UserId    int
	Data      string
	TTL       time.Duration
	CreatedAt time.Time
}
