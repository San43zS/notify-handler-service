package main

import (
	"Notify-handler-service/internal/app"
	"context"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	server, err := app.New()
	if err != nil {
		log.Fatalf("error occured while creating server: %v", err)
	}

	if err := server.Start(context.Background()); err != nil {
		log.Fatalf("error occured while running http server: %v", err)
	}
}
