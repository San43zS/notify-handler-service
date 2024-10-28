package msg

import "time"

type MSG struct {
	UserId    int           `json:"user_id"`
	Data      string        `json:"data"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
}
