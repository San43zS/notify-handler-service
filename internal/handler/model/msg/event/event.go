package event

import "time"

const (
	AddNotify        string = "add-notify"
	SendNotify       string = "send-notify"
	GetCurrentNotify string = "get-current-notify"
	ChangeExpired    string = "change-expired"
	GetOldNotify     string = "get-old-notify"
)

const (
	TTL     time.Duration = 15 * time.Second
	User_ID int           = 15
)
