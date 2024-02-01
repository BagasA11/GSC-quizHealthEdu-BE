package dto

// quiz id
// {
// 	question: [{id, question, Answer, img, opt: [{id, alpha, txt, color}]}]
// 	question: [{id, question, Answer, img, opt: [{id, alpha, txt, color}]}]
// 	question: [{id, question, Answer, img, opt: [{id, alpha, txt, color}]}]
// 	question: [{id, question, Answer, img, opt: [{id, alpha, txt, color}]}]
// }

/*answering and scoring Quiz*/
type Answer struct {
	QuizID uint `json:"quizID" binding:"required"`
	Quest  []Question
}

type Question struct {
	ID        uint   `json:"questID" binding:"required"`
	Quest     string `json:"question" binding:"required"`
	Answer    string `json:"answ" binding:"required"`
	Img       string `json:"img"`
	PromptOpt string `json:"prompt" binding:"required"`
	Opt       []Options
}

type Options struct {
	ID    uint   `json:"questID" binding:"required"`
	Alpha string `json:"alpha" binding:"required"`
	Txt   string `json:"detail" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type Scoring struct {
	QuizID uint `json:"quizID" binding:"required"`
	Quest  []QuesScoring
}
type QuesScoring struct {
	ID     uint   `json:"ID" binding:"required"`
	Ans    string `json:"QuestAnsw" binding:"required"`
	Choose string `json:"answ" binding:"required"`
}
