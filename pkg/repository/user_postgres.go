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

func (r *UserPostgres) GetById(userId int) (survey.User, error) {
	var user survey.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}
