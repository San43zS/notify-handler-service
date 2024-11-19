package msg

import "time"

const (
	CurrentStatus = "not_viewed"
	OldStatus     = "viewed"
)

type MSG struct {
	UserId    int           `json:"user_id"`
	Content   string        `json:"content"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

type Notify struct {
	Id        string        `json:"id"`
	UserId    int           `json:"user_id"`
	Status    string        `json:"status"`
	Content   string        `json:"content"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

type Expired struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Message struct {
	Type string `json:"type"`
	Data MSG    `json:"data"`
}

type Common struct {
	Type string `json:"type"`
}
