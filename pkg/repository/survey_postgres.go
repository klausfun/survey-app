package repository

import (
	"errors"
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

	if survey.Types == "free" {
		return surveyId, tx.Commit()
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

func (r *SurveyPostgres) Update(userId, surveyId int, input survey.UpdateSurveyInput) error {
	// получаем тип
	var types survey.Types
	queryType := fmt.Sprintf("SELECT types FROM %s WHERE id = $1", surveysTable)
	err := r.db.Get(&types, queryType, surveyId)
	if err != nil {
		return err
	}

	var answersId []int
	queryId := fmt.Sprintf("SELECT ans.id FROM %s ans"+
		" INNER JOIN %s qst_ans on ans.id = qst_ans.answer_id"+
		" INNER JOIN %s qst 	on qst.id = qst_ans.question_id"+
		" INNER JOIN %s sur_qst on qst.id = sur_qst.question_id"+
		" INNER JOIN %s sur on sur.id = sur_qst.survey_id"+
		" WHERE sur.id = $1", answersTable, questionsAnswersTable, questionsTable, surveysQuestionsTable, surveysTable)
	err = r.db.Select(&answersId, queryId, surveyId)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if types.Types == "multiple" {
		err = r.idVerification(input, answersId)
		if err != nil {
			return err
		}

		for _, id := range input.Response {
			idFloat, ok := id.(float64)
			if !ok {
				return errors.New("response id is of invalid type")
			}
			idInt := int(idFloat)

			ok, err := r.votedValidate(userId, idInt, surveyId, "multiple")
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("This user has already voted!")
			}

			queryAnswer := fmt.Sprintf("UPDATE %s SET amount = amount + 1"+
				" WHERE id = $1", answersTable)
			_, err = tx.Exec(queryAnswer, idInt)
			if err != nil {
				tx.Rollback()
				return err
			}

			var id int
			queryVotes := fmt.Sprintf("INSERT INTO %s (user_id, answer_id, survey_id) VALUES ($1, $2, $3) RETURNING id", votesTable)
			row := tx.QueryRow(queryVotes, userId, idInt, surveyId)
			if err := row.Scan(&id); err != nil {
				tx.Rollback()
				return err
			}

		}
		return tx.Commit()

	} else if types.Types == "single" {
		err = r.idVerification(input, answersId)
		if err != nil {
			return err
		}

		if len(input.Response) > 1 {
			return errors.New("multiple choice is not available in a single survey")
		}

		idFloat, ok := input.Response[0].(float64)
		if !ok {
			return errors.New("response id is of invalid type")
		}
		idInt := int(idFloat)

		ok, err := r.votedValidate(userId, idInt, surveyId, "single")
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("this user has already voted")
		}

		tx, err := r.db.Begin()
		if err != nil {
			return err
		}

		queryAnswer := fmt.Sprintf("UPDATE %s SET amount = amount + 1"+
			" WHERE id = $1", answersTable)
		_, err = tx.Exec(queryAnswer, idInt)
		if err != nil {
			tx.Rollback()
			return err
		}

		var id int
		queryVotes := fmt.Sprintf("INSERT INTO %s (user_id, answer_id, survey_id) VALUES ($1, $2, $3) RETURNING id", votesTable)
		row := tx.QueryRow(queryVotes, userId, idInt, surveyId)
		if err := row.Scan(&id); err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit()

	} else if types.Types == "free" {
		if len(input.Response) > 1 {
			return errors.New("A multiple response in a free survey is not available")
		}

		tx, err := r.db.Begin()
		if err != nil {
			return err
		}

		var answerId int
		createAnswersQuery := fmt.Sprintf("INSERT INTO %s (description, amount) VALUES ($1, 1) RETURNING id", answersTable)
		row := tx.QueryRow(createAnswersQuery, input.Response[0])
		if err := row.Scan(&answerId); err != nil {
			tx.Rollback()
			return err
		}

		questionId := surveyId // пока не изменю логику в бд
		createQuestionsAnswersQuery := fmt.Sprintf("INSERT INTO %s (answer_id, question_id) VALUES ($1, $2)", questionsAnswersTable)
		_, err = tx.Exec(createQuestionsAnswersQuery, answerId, questionId)
		if err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit()
	}

	return errors.New("non-existing survey type")
}

func (r *SurveyPostgres) votedValidate(userId, answerId, surveyId int, types string) (bool, error) {
	if types == "multiple" {
		var vote []survey.Vote
		query := fmt.Sprintf("SELECT user_id, answer_id"+
			" FROM %s WHERE user_id = $1 AND answer_id = $2", votesTable)
		err := r.db.Select(&vote, query, userId, answerId)
		if err != nil {
			return false, err
		}

		if vote == nil {
			return true, nil
		}

		return false, nil
	} else if types == "single" {
		var vote []survey.Vote
		query := fmt.Sprintf("SELECT user_id, survey_id"+
			" FROM %s WHERE user_id = $1 AND survey_id = $2", votesTable)
		err := r.db.Select(&vote, query, userId, surveyId)
		if err != nil {
			return false, err
		}

		if vote == nil {
			return true, nil
		}

		return false, nil
	}

	return false, errors.New("non-existent type of survey")
}

func (r *SurveyPostgres) idVerification(input survey.UpdateSurveyInput, answersId []int) error {
	for _, inputId := range input.Response {
		flag := 0
		for _, answerId := range answersId {
			if int(inputId.(float64)) == answerId {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return errors.New("invalid id for this survey")
		}
	}

	return nil
}
