package controllers

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"
	"BagasA11/GSC-quizHealthEdu-BE/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

/*Create User or Player service*/
func (uc *UserController) CreateUser(c *gin.Context) {
	req := new(dto.UserCreate)
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
	err = uc.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *UserController) CreateAdmin(c *gin.Context) {
	//get Request data
	req := new(dto.AdminCreate)
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
	//verify token type
	tokenType, exist := c.Get("TokenType")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "TokenType value not set",
		})
		return
	}

	if tokenType.(string) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User are not allowed to create Admin",
		})
		return
	}
	//service
	err = uc.service.CreateAdmin(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create admin",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *UserController) UpdateUsername(ctx *gin.Context) {
	req := new(dto.UpdateUsername)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		ctx.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	//Get User ID from Token
	ID, exist := ctx.Get("ID")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "id value not found",
		})
		return
	}
	//verify token type
	tokenType, exist := ctx.Get("TokenType")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "TokenType field not set",
		})
		return
	}

	if tokenType.(string) == "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Admin are not allowed change username",
		})
		return
	}
	err = uc.service.UpdateUsername(ID.(uint), *req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update password",
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	req := new(dto.UpdatePassword)
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
	//Jwt claims
	rawID, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "id value not found",
		})
		return
	}
	ID, _ := rawID.(uint)
	// service
	err = uc.service.UpdatePassword(ID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update password",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *UserController) Delete(c *gin.Context) {
	req := new(*dto.DeleteUser)
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

	id, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "id value not found",
		})
		return
	}
	err = uc.service.DeleteUser(id.(uint), *req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update password",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success to Delete Account",
	})
}
