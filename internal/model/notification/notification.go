package notification

import "time"

type Notification struct {
	UserId    int
	Data      string
	TTL       time.Duration
	CreatedAt time.Time
}

type Notify struct {
	Number    int       `json:"number"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}
