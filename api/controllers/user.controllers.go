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

}

func (uc *UserController) CreateAdmin(c *gin.Context) {

}
