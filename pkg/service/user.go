package service

import (
	survey "survey_app"
	"survey_app/pkg/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]survey.User, error) {
	return s.repo.GetAll()
}
