package survey

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

//type Surveys struct {
//	Id    int    `json:"id" db:"id"`
//	Types string `json:"types" db:"types"`
//}

//
//type UsersSurveys struct {
//	Id       int
//	UserId   int
//	SurveyId int
//}
//
//type Questions struct {
//	Id                  int    `json:"id"`
//	QuestionDescription string `json:"question_description" binding:"required"`
//}
//
//type SurveysQuestions struct {
//	Id         int
//	QuestionId int
//	SurveyId   int
//}
//
//type Answers struct {
//	Id                 int    `json:"id"`
//	AnswersDescription string `json:"answers_description" binding:"required"`
//	Amount             int    `json:"amount"`
//}
//
//type QuestionsAnswers struct {
//	Id         int
//	QuestionId int
//	AnswerId   int
//}
