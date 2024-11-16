package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	ChannelName = "Notify"
)

type Config struct {
	URL      string
	Password string
	Username string
}

func NewConfig() Config {
	host := viper.GetString("db.redis.host")
	port := viper.GetString("db.redis.port")

	return Config{
		URL:      fmt.Sprintf("%s:%s", host, port),
		Password: viper.GetString("db.redis.password"),
		Username: viper.GetString("db.redis.username"),
	}
}
