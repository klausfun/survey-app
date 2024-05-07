package service

import (
	survey "survey_app"
	"survey_app/pkg/repository"
)

type Authorization interface {
	CreateUser(user survey.User) (int, error)
	GenerateToken(email, password, role string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Surveys interface {
	CreateSurvey(userId int, survey survey.Data) (int, error)
	GetAll(userId int) ([]survey.Surveys, error)
	GetById(userId, surveyId int) (survey.Surveys, error)
	Delete(userId, surveyId int) error
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
		Surveys:       NewSurveyService(repos.Surveys),
	}
}
