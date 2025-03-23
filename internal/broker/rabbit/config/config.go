package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	ConsumerQueueName    = "yellowQueue"
	ConsumerExchangeName = "test"
	ConsumerRoutingKey   = "yellow"
)

const (
	HandlerServiceConsumer = "HandlerServiceConsumer"
	ProducerQueueName      = "redQueue"
	ProducerExchangeName   = "test"
	ProducerRoutingKey     = "redQueue"
)

type Config struct {
	URL    string
	Driver string
}

type amqpParams struct {
	host     string
	port     string
	user     string
	password string
}

func getAMQPParams() *amqpParams {
	return &amqpParams{
		host:     viper.GetString("AMQP.HOST"),
		port:     viper.GetString("AMQP.PORT"),
		user:     viper.GetString("AMQP.USER"),
		password: viper.GetString("AMQP.PASSWORD"),
	}
}

func (amqp amqpParams) ParseURL() string {
	template := viper.GetString("AMQP.URLTEMPLATE")

	return fmt.Sprintf(template, amqp.user, amqp.password, amqp.host, amqp.port)
}

func NewConfig() *Config {
	return &Config{
		URL:    getAMQPParams().ParseURL(),
		Driver: viper.GetString("AMQP.DRIVER"),
	}
}
