package redis

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	ChannelName = "__keyevent@0__:expired"
)

type Config struct {
	URL      string
	Password string
	Username string
}

func NewConfig() Config {
	host := viper.GetString("DB.REDIS.HOST")
	port := viper.GetString("DB.REDIS.PORT")

	return Config{
		URL:      fmt.Sprintf("%s:%s", host, port),
		Password: viper.GetString("DB.REDIS.PASSWORD"),
		Username: viper.GetString("DB.REDIS.USERNAME"),
	}
}
