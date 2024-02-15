package controllers

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/service"
	"net/http"
	"strconv"

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

func (ss *ScoreController) Rank(c *gin.Context) {
	//token validation
	_, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusBadRequest, "user id not found")
		return
	}
	quizID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "cannot convert string to integer or url id is unavailable",
			"error":   err,
		})
		return
	}
	score, err := ss.service.Rank(uint(quizID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "cannot serve request",
			"error":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": score,
	})
}

func (sc *ScoreController) History(c *gin.Context) {
	//get user id from token
	userID, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusBadRequest, "user id not found")
		return
	}

	s, err := sc.service.GetHistory(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "cannot serve request",
			"error":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": s,
	})
}
