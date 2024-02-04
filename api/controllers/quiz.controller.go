package controllers

// 	"BagasA11/GSC-quizHealthEdu-BE/api/service"

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"
	"BagasA11/GSC-quizHealthEdu-BE/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type QuizController struct {
	service *service.QuizService
}

func NewQuizController() *QuizController {
	return &QuizController{
		service: service.NewQuizService(),
	}
}

func (qc *QuizController) Create(c *gin.Context) {
	req := new(dto.QuizCreate)
	err := c.ShouldBindJSON(&req)
	if err != nil {
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
	_, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token id not found",
		})
		return
	}
	tokenTyp, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "TokenType value not set",
		})
		return
	}
	if tokenTyp.(string) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User are not allowed to create Quiz",
		})
		return
	}
	err = qc.service.CreateQuiz(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create Quiz",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (qc *QuizController) FindCheapest(c *gin.Context) {
	data, err := qc.service.Cheapest()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "ok",
		"data":    data,
	})
}
