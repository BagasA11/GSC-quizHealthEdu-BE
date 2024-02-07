package dto

import "mime/multipart"

type QuizCreate struct {
	Title string                `form:"title"`
	Topic string                `form:"topic"`
	Desc  string                `form:"desc"`
	Img   *multipart.FileHeader `form:"file"`
}

type Question struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type Option struct {
	Alphabet string  `json:"options" binding:"required,max=1"`
	Txt      string  `json:"desc" binding:"required"`
	Color    *string `json:"color" binding:"required"`
}

type FindTopic struct {
	Topic string `json:"topic" binding:"required"`
}

type FindTitle struct {
	Title string `json:"title" binding:"required"`
}
