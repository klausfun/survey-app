package survey

type Data struct {
	Types               string        `json:"types" binding:"required"`
	AnswersDescription  []interface{} `json:"answers_description" binding:"required"`
	QuestionDescription string        `json:"question_description" binding:"required"`
	CountAnswers        int           `json:"countAnswers" binding:"required"`
}

//
//type Surveys struct {
//	Id    int    `json:"id"`
//	Types string `json:"types" binding:"required"`
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
