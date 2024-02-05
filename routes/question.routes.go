package routes

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/controllers"
	"BagasA11/GSC-quizHealthEdu-BE/middleware"

	"github.com/gin-gonic/gin"
)

func QuestionRoutes(group *gin.RouterGroup) {
	qc := controllers.NewQuestionController()
	//create new quiz
	group.POST("/question/:quizid/create", middleware.JwtAuth(), qc.Create)
	group.POST("/question/update/:id", middleware.JwtAuth(), qc.Edit)
}
