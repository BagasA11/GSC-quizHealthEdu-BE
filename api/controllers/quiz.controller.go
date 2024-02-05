package controllers

// 	"BagasA11/GSC-quizHealthEdu-BE/api/service"

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"
	"BagasA11/GSC-quizHealthEdu-BE/api/service"
	"fmt"
	"net/http"
	"strconv"

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

func (qc *QuizController) FindTopic(c *gin.Context) {
	req := new(dto.FindTopic)
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
	q, err := qc.service.FindTopic(req.Topic)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"massage": "topic not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "topic found",
		"data":    q,
	})
}

func (qc *QuizController) FindTitle(c *gin.Context) {
	req := new(dto.FindTitle)
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
	q, err := qc.service.FindTopic(req.Title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"massage": "title not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "title found",
		"data":    q,
	})
}

func (qc *QuizController) FindID(c *gin.Context) {
	//retrieve id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	data, err := qc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"massage": "Quiz id not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "success",
		"data":    data,
	})
}

func (qc *QuizController) Edit(c *gin.Context) {
	//token validation
	//get token type
	tokenType, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusBadRequest, "token type not set")
		return
	}
	if tokenType.(string) != "admin" {
		c.JSON(http.StatusUnauthorized, "you are not allowed to modificate quiz entity")
		return
	}
	//parsing request body
	req := new(dto.QuizCreate)
	err := c.ShouldBindJSON(&req)
	//request validation
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	//serving to update data
	err = qc.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Failed to update Quiz",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "success",
	})
}

func (qc *QuizController) Delete(c *gin.Context) {
	//token validation
	//get token type
	tokenType, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusBadRequest, "token type not set")
		return
	}
	if tokenType.(string) != "admin" {
		c.JSON(http.StatusUnauthorized, "user are not allowed to delete quiz entity")
		return
	}
	//retrieve id from url parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = qc.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Failed to delete Quiz",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "success to delete quiz object",
	})
}
