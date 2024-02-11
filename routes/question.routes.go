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
	group.GET("/quiz/:id/question", middleware.JwtAuth(), qc.GetQuestionAndOption)
	//create new quiz
	group.POST("/quiz/:id/question/create", middleware.JwtAuth(), qc.Create)
	group.PUT("/question/update/:id", middleware.JwtAuth(), qc.Edit)
	group.PUT("/question/upload/:id", middleware.JwtAuth(), qc.UploadFile)
	group.DELETE("/question/delete/:id", middleware.JwtAuth(), qc.Delete)
}
