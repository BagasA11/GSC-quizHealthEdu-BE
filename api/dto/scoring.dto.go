package dto

type Score struct {
	QuizID uint    `json:"quizId" binding:"required"`
	UserID uint    `json:"userId" binding:"required"`
	Score  float64 `json:"score" binding:"required"`
}

type Answer struct {
	QuestID  uint   `json:"questID" binding:"required"`
	Answer   string `json:"ans" binding:"required,max=1"`
	Checkbox string `json:"checkbox" binding:"required,max=1"`
}
