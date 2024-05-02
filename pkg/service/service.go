package service

import (
	survey "survey_app"
	"survey_app/pkg/repository"
)

type Authorization interface {
	CreateUser(user survey.User) (int, error)
	GenerateToken(email, password, role string) (string, error)
	ParseToken(token string) (int, error)
}

type Surveys interface {
}

type Users interface {
}

type Service struct {
	Authorization
	Surveys
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
