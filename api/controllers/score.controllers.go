package controllers

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"

	"BagasA11/GSC-quizHealthEdu-BE/api/service"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
)

// "BagasA11/GSC-quizHealthEdu-BE/api/service"
// "strconv"
// "fmt"
// "github.com/joho/godotenv"
// "BagasA11/GSC-quizHealthEdu-BE/api/dto"
// os

var Store = sessions.NewCookieStore([]byte("my-sessions-screet"))

type ScoreController struct {
	service *service.ScoreService
}

func NewScoreController() *ScoreController {
	return &ScoreController{
		service: service.NewScoreService(),
	}
}

func (sc *ScoreController) CreateOrUpdate(c *gin.Context) {

	//session object instance
	session, err := Store.Get(c.Request, "attempt-quiz")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//token validation
	//get user id from header
	/*
		userID, exist := c.Get("ID")
		if !exist {
			c.JSON(400, gin.H{
				"error": "user id is not exist",
			})
		}
		//get quiz id from url
		quizID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{
				"massage": "id on url not found or Failed to convert string into integer",
				"error":   err.Error(),
			})
			return
		}
	*/

	req := new(dto.Answer)
	// request validation
	if err = c.ShouldBindJSON(&req); err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	//length validation
	if req.Num > req.Length {
		c.JSON(400, gin.H{
			"massage": "num is overide max number",
		})
		return
	}
	if req.Num <= 0 {
		c.JSON(400, gin.H{
			"error": "Num must be greater than 0",
		})
	}
	var point float32 = 0
	// var i float32 = float32(1)/float32(int32(5)) * 100

	if req.Num == 1 {

		if req.Answer == req.Checkbox {
			point = float32(1) / float32(int32(req.Length)) * 100
		}
		session.Values["point"] = point
		err := session.Save(c.Request, c.Writer)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, point)
		return
	}

	if req.Num == req.Length {
		if req.Answer == req.Checkbox {
			point = float32(1) / float32(int32(req.Length)) * 100
		}

		if session.Values["point"] != nil {
			point += session.Values["point"].(float32)
		}

		session.Options.MaxAge = -1
		err := session.Save(c.Request, c.Writer)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, "finish")
		return
	}

	if req.Answer == req.Checkbox {
		point = float32(1) / float32(int32(req.Length)) * 100
	}
	if session.Values["point"] != nil {
		point += session.Values["point"].(float32)
	}
	session.Values["point"] = point
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"point":  point,
		"result": float32(1) / float32(int32(req.Length)) * 100,
	})
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
