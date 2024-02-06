package routes

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/controllers"
	"BagasA11/GSC-quizHealthEdu-BE/middleware"

	"github.com/gin-gonic/gin"
)

func QuestionRoutes(group *gin.RouterGroup) {
	qc := controllers.NewQuestionController()
	//get question detail by id
	group.GET("/question/:id", middleware.JwtAuth(), qc.FindID)
	group.GET("/question/:quizid?page", middleware.JwtAuth(), qc.GetQuestionAndOption)
	//create new quiz
	group.POST("/question/:quizid/create", middleware.JwtAuth(), qc.Create)
	group.PUT("/question/update/:id", middleware.JwtAuth(), qc.Edit)
	group.DELETE("/question/delete/:id", middleware.JwtAuth(), qc.Delete)
}
