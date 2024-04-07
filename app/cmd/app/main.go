package main

import (
	"api/internal/controller"
	"api/internal/repository"
	"api/internal/service"
	"api/pkg/postgres"
	"api/pkg/server"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	pg, err := postgres.New(context.Background(), "host=postgres dbname=app_db user=postgres sslmode=disable")
	if err != nil {
		logrus.Fatalf("Unable to connection to database: %v\n", err)
	}
	repos := repository.NewRepository(pg)
	services := service.NewService(repos)
	router := controller.NewRouter(services)
	srv := server.New(viper.GetString("port"), router)

	if err := srv.Run(); err != nil {
		logrus.Fatal(err)
	}
}
