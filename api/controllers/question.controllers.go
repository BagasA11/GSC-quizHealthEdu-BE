package controllers

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"
	"BagasA11/GSC-quizHealthEdu-BE/api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type QuestionController struct {
	service *service.QuestionService
}

func NewQuestionController() *QuestionController {
	return &QuestionController{
		service: service.NewQuestionService(),
	}
}

func (qc *QuestionController) Create(c *gin.Context) {
	//token validation
	tkType, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "TokenType value not set",
		})
		return
	}
	if tkType.(string) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User are not allowed to create Question",
		})
		return
	}
	//get quiz id from url
	quizId, err := strconv.Atoi(c.Param("quizid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	//get request data
	req := new(dto.Question)
	err = c.ShouldBindJSON(&req)
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
	err = qc.service.Create(uint(quizId), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create Question",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (qc *QuestionController) FindID(c *gin.Context) {
	//token validation
	_, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusBadRequest, "token not found")
		return
	}
	//retrieve id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	q, err := qc.service.FindID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"massages": "question id not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "question was found",
		"data":    q,
	})
}

func (qc *QuestionController) GetQuestionAndOption(c *gin.Context) {
	_, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "token not available",
		})
		return
	}
	//retrieve id from url
	// url/quizid
	quizID, err := strconv.Atoi(c.Param("quizid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "quiz id on url not found")
		return
	}

	q, err := qc.service.GetQuizAndOption(uint(quizID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "invalid request",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": q,
	})
}

func (qc *QuestionController) Edit(c *gin.Context) {
	tkType, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "TokenType value not set",
		})
		return
	}
	if tkType.(string) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User are not allowed to modify Question",
		})
		return
	}
	//get quiz id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	//get request data
	//get request data
	req := new(dto.Question)
	err = c.ShouldBindJSON(&req)
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
	err = qc.service.Updates(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update Question",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (qc *QuestionController) Delete(c *gin.Context) {
	//token validation
	typ, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "token not found",
		})
		return
	}
	if typ.(string) != "admin" {
		c.JSON(http.StatusForbidden, "user not allowed to delete question")
		return
	}
	//get id from url
	//url/question/id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "invalid request",
			"error":   err,
		})
		return
	}
	//serving delete behavior
	err = qc.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "delete question success")
}
