package survey

type Surveys struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Done     bool   `json:"done"`
	WinnerId int    `json:"winnerId"`
}

type UsersSurveys struct {
	Id       int
	UserId   int
	SurveyId int
}

type Questions struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type SurveysQuestions struct {
	Id         int
	QuestionId int
	SurveyId   int
}

type Answers struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type QuestionsAnswers struct {
	Id         int
	QuestionId int
	AnswerId   int
}
