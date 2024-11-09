package event

import "time"

const (
	AddNotify  string = "add-notify"
	SendNotify string = "send-notify"
)

const (
	TTL     time.Duration = 15 * time.Second
	User_ID int           = 15
)
