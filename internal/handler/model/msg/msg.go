package msg

import "time"

const (
	CurrentStatus = "not_viewed"
	OldStatus     = "viewed"
)

type MSG struct {
	Type      string        `json:"type"`
	UserId    int           `json:"user_id"`
	Content   []byte        `json:"content"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

type Expired struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}
