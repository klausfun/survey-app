package main

import (
	"github.com/sirupsen/logrus"
	survey "survey_app"
	"survey_app/pkg/handler"
	"survey_app/pkg/repository"
	"survey_app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(survey.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
