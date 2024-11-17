package encoding

import (
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/exp/rand"
	"strconv"
	"time"
)

type Service interface {
	NotificationID(userId int) string
}

type service struct {
	randomNumber int
}

func New() Service {
	return &service{
		randomNumber: rand.Int(),
	}
}

func (s service) NotificationID(userId int) string {
	t := time.Now()

	hash := sha256.Sum256([]byte(strconv.Itoa(s.randomNumber) + strconv.Itoa(userId) + t.String()))
	return base64.StdEncoding.EncodeToString(hash[:])
}
