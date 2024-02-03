package dto

type QuizCreate struct {
	Title string `json:"title" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}
