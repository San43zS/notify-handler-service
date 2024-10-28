package app

import (
	"Notify-handler-service/internal/server"
	"Notify-handler-service/internal/service"
	"Notify-handler-service/internal/storage/config"
	"Notify-handler-service/internal/storage/db/redis"
	"context"
	"log"
)

type App struct {
	server service.Service
}

func New() (*App, error) {
	storage, err := redis.New(config.NewConfig())
	if err != nil {
		log.Fatal("Error while connecting to redis: ", err)
		return &App{}, err
	}

	srv := service.New(storage)

	app := &App{
		server: srv,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.server)
	if err != nil {
		return err
	}

	//TODO:
	//if err := srv.Start(ctx); err != nil {
	//	return err
	//}
	log.Println("Server stopped")
	return nil
}
