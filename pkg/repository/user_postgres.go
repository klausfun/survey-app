package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	survey "survey_app"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAll() ([]survey.User, error) {
	var users []survey.User
	query := fmt.Sprintf("SELECT * FROM %s", userTable)
	err := r.db.Select(&users, query)

	return users, err
}
