package dto

type QuizCreate struct {
	Title string `json:"title" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

type Question struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type Option struct {
	Alphabet string `json:"options" binding:"required,max=1"`
	Txt      string `json:"desc" binding:"required"`
	Color    string `json:"color" binding:"required"`
}
