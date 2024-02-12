package controllers

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/service"

	"github.com/gin-gonic/gin"
)

type ScoreController struct {
	service *service.ScoreService
}

func NewScoreController() *ScoreController {
	return &ScoreController{
		service: service.NewScoreService(),
	}
}

func (sc *ScoreController) Create(c *gin.Context) {

}
