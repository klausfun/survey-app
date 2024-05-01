package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
