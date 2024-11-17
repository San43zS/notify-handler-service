package notification

import "time"

type Notification struct {
	Id        string
	UserId    int
	Status    string
	Data      string
	TTL       time.Duration
	CreatedAt time.Time
	ExpiredAt time.Time
}

type Notify struct {
	Number    int       `json:"number"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}
