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
	}
}
