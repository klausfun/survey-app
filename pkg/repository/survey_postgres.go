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

func (r *SurveyPostgres) GetAll(userId int) ([]survey.Surveys, error) {
	var surveys []survey.Surveys

	query := fmt.Sprintf("SELECT sur.id, sur.types,"+
		" qst.description as question_description"+
		" FROM %s sur"+
		" INNER JOIN %s us 	    on us.survey_id = sur.id"+
		" INNER JOIN %s sur_qst on sur.id = sur_qst.survey_id"+
		" INNER JOIN %s qst 	on sur_qst.question_id = qst.id"+
		" WHERE us.user_id = $1",
		surveysTable, usersSurveysTable, surveysQuestionsTable, questionsTable)
	err := r.db.Select(&surveys, query, userId)

	for i, curSurvey := range surveys {
		var answers []survey.Answers

		queryAnswers := fmt.Sprintf("SELECT ans.description as description, ans.amount as amount FROM %s sur"+
			" INNER JOIN %s us      on us.survey_id = sur.id"+
			" INNER JOIN %s sur_qst on sur.id = sur_qst.survey_id"+
			" INNER JOIN %s qst 	on sur_qst.question_id = qst.id"+
			" INNER JOIN %s qst_ans on qst_ans.question_id = qst.id"+
			" INNER JOIN %s ans 	on ans.id = qst_ans.answer_id"+
			" WHERE us.user_id = $1 AND sur.id = $2",
			surveysTable, usersSurveysTable, surveysQuestionsTable, questionsTable, questionsAnswersTable, answersTable)
		err := r.db.Select(&answers, queryAnswers, userId, curSurvey.Id)
		if err != nil {
			return nil, err
		}

		surveys[i].AnswersDescription = answers
	}

	return surveys, err
}

func (r *SurveyPostgres) GetById(userId, surveyId int) (survey.Surveys, error) {
	var sur survey.Surveys
	var answers []survey.Answers

	queryAnswers := fmt.Sprintf("SELECT ans.description as description, ans.amount as amount FROM %s sur"+
		" INNER JOIN %s us      on us.survey_id = sur.id"+
		" INNER JOIN %s sur_qst on sur.id = sur_qst.survey_id"+
		" INNER JOIN %s qst 	on sur_qst.question_id = qst.id"+
		" INNER JOIN %s qst_ans on qst_ans.question_id = qst.id"+
		" INNER JOIN %s ans 	on ans.id = qst_ans.answer_id"+
		" WHERE us.user_id = $1 AND sur.id = $2",
		surveysTable, usersSurveysTable, surveysQuestionsTable, questionsTable, questionsAnswersTable, answersTable)
	err := r.db.Select(&answers, queryAnswers, userId, surveyId)
	if err != nil {
		return sur, err
	}

	query := fmt.Sprintf("SELECT sur.id, sur.types,"+
		" qst.description as question_description"+
		" FROM %s sur"+
		" INNER JOIN %s us 	    on us.survey_id = sur.id"+
		" INNER JOIN %s sur_qst on sur.id = sur_qst.survey_id"+
		" INNER JOIN %s qst 	on sur_qst.question_id = qst.id"+
		" WHERE us.user_id = $1 AND sur.id = $2",
		surveysTable, usersSurveysTable, surveysQuestionsTable, questionsTable)
	err = r.db.Get(&sur, query, userId, surveyId)
	sur.AnswersDescription = answers

	return sur, err
}

func (r *SurveyPostgres) Delete(userId, surveyId int) error {
	query := fmt.Sprintf("DELETE FROM %s sur USING %s us WHERE sur.id = us.survey_id "+
		"AND us.user_id = $1 AND us.survey_id = $2", surveysTable, usersSurveysTable)
	_, err := r.db.Exec(query, userId, surveyId)

	return err
}
