package dto

type Score struct {
	QuizID uint    `json:"quizId" binding:"required"`
	UserID uint    `json:"userId" binding:"required"`
	Score  float64 `json:"score" binding:"required"`
}
