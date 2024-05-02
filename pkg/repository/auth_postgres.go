package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	survey "survey_app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user survey.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email, role)"+
		" VALUES ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password, role string) (survey.User, error) {
	var user survey.User
	query := fmt.Sprintf("SELECT id FROM %s "+
		"WHERE email = $1 AND password_hash = $2 AND role = $3", userTable)
	err := r.db.Get(&user, query, email, password, role)

	return user, err
}
