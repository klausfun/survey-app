package survey

import "errors"

type Data struct {
	Id                  int           `json:"id" db:"id"`
	Types               string        `json:"types" binding:"required" db:"types"`
	AnswersDescription  []interface{} `json:"answers_description" binding:"required" db:"answers_description"`
	QuestionDescription string        `json:"question_description" binding:"required" db:"question_description"`
}

type Surveys struct {
	Id                  int       `json:"id" db:"id"`
	Types               string    `json:"types" db:"types"`
	QuestionDescription string    `json:"question_description" db:"question_description"`
	AnswersDescription  []Answers `json:"answers_description" db:"answers_description"`
}

type Answers struct {
	Description string `json:"description" db:"description"`
	Amount      int    `json:"amount" db:"amount"`
}

type Vote struct {
	UserId   int `json:"userId" db:"user_id"`
	AnswerId int `json:"answerId" db:"answer_id"`
	SurveyId int `json:"surveyId" db:"survey_id"`
}

type Types struct {
	Types string `json:"types"`
}

type UpdateSurveyInput struct {
	Response []interface{} `json:"response"`
}

func (i UpdateSurveyInput) Validate() error {
	if i.Response == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
