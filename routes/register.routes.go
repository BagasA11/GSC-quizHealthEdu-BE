package routes

//register all routes
import (
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	//endpoint to handle user
	UserRoutes(apiGroup)
	//endpoint to handle quiz
	QuizRoutes(apiGroup)
	r.Run(":" + os.Getenv("HOST_PORT"))
}
