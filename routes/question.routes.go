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
	group.GET("/question/quest-option/:quizid", middleware.JwtAuth(), qc.GetQuestionAndOption)
	//create new quiz
	group.POST("/question/create/:quizid", middleware.JwtAuth(), qc.Create)
	group.PUT("/question/update/:id", middleware.JwtAuth(), qc.Edit)
	group.DELETE("/question/delete/:id", middleware.JwtAuth(), qc.Delete)
}
