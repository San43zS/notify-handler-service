package msg

import "time"

type MSG struct {
	Type string `json:"type"`

	Content Data
}

type Data struct {
	Data []byte `json:"data"`
}

type Message struct {
	UserId    int
	CreatedAt time.Time `json:"created_at"`
	Data      Data
}
