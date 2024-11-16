package app

import (
	"Notify-handler-service/internal/broker"
	"Notify-handler-service/internal/server"
	"Notify-handler-service/internal/service"
	storage "Notify-handler-service/internal/storage"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"fmt"
	"log"
)

type App struct {
	server  service.Service
	storage redis.Store
	broker  broker.Broker
}

func New() (*App, error) {
	strg, err := storage.New()
	if err != nil {
		log.Fatal("Error while connecting to redis: ", err)
		return &App{}, err
	}
	//////////////????????????????????????????>
	broker, err := broker.New()
	if err != nil {
		return &App{}, err
	}

	srv := service.New(strg, broker)

	app := &App{
		server:  srv,
		storage: strg,
		broker:  broker,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.server, a.storage.PubSub(), a.broker)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	log.Println("Server stopped")

	return nil
}
