package repository

import (
	"github.com/jmoiron/sqlx"
	survey "survey_app"
)

type Authorization interface {
	CreateUser(user survey.User) (int, error)
	GetUser(email, password, role string) (survey.User, error)
}

type Surveys interface {
	CreateSurvey(userId int, survey survey.Data) (int, error)
	GetAll(userId int) ([]survey.Surveys, error)
	GetById(userId, surveyId int) (survey.Surveys, error)
	Delete(userId, surveyId int) error
}

type Users interface {
}

type Repository struct {
	Authorization
	Surveys
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Surveys:       NewSurveyPostgres(db),
	}
}
