package service

import (
	survey "survey_app"
	"survey_app/pkg/repository"
)

type SurveyService struct {
	repo repository.Surveys
}

func NewSurveyService(repo repository.Surveys) *SurveyService {
	return &SurveyService{repo: repo}
}

func (s *SurveyService) CreateSurvey(userId int, survey survey.Data) (int, error) {
	return s.repo.CreateSurvey(userId, survey)
}

func (s *SurveyService) GetAll(userId int) ([]survey.Surveys, error) {
	return s.repo.GetAll(userId)
}