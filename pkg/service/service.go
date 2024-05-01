package service

import "survey_app/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
