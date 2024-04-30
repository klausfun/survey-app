package main

import (
	"github.com/sirupsen/logrus"
	survey "survey_app"
)

func main() {
	srv := new(survey.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
