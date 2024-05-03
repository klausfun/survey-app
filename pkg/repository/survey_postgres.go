package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	survey "survey_app"
)

type SurveyPostgres struct {
	db *sqlx.DB
}

func NewSurveyPostgres(db *sqlx.DB) *SurveyPostgres {
	return &SurveyPostgres{db: db}
}

func (r *SurveyPostgres) CreateSurvey(userId int, survey survey.Data) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var surveyId int
	createSurveyQuery := fmt.Sprintf("INSERT INTO %s (types) VALUES ($1) RETURNING id", surveysTable)
	row := tx.QueryRow(createSurveyQuery, survey.Types)
	if err := row.Scan(&surveyId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersSurveysQuery := fmt.Sprintf("INSERT INTO %s (user_id, survey_id) VALUES ($1, $2)", usersSurveysTable)
	_, err = tx.Exec(createUsersSurveysQuery, userId, surveyId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var questionId int
	createQuestionsQuery := fmt.Sprintf("INSERT INTO %s (description) VALUES ($1) RETURNING id", questionsTable)
	row = tx.QueryRow(createQuestionsQuery, survey.QuestionDescription)
	if err := row.Scan(&questionId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createSurveysQuestionsQuery := fmt.Sprintf("INSERT INTO %s (survey_id, question_id) VALUES ($1, $2)", surveysQuestionsTable)
	_, err = tx.Exec(createSurveysQuestionsQuery, surveyId, questionId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var answerId int
	for _, temp := range survey.AnswersDescription {
		createAnswersQuery := fmt.Sprintf("INSERT INTO %s (description) VALUES ($1) RETURNING id", answersTable)
		row = tx.QueryRow(createAnswersQuery, temp)
		if err := row.Scan(&answerId); err != nil {
			tx.Rollback()
			return 0, err
		}

		createQuestionsAnswersQuery := fmt.Sprintf("INSERT INTO %s (answer_id, question_id) VALUES ($1, $2)", questionsAnswersTable)
		_, err = tx.Exec(createQuestionsAnswersQuery, answerId, questionId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return surveyId, tx.Commit()
}
