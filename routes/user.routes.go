package routes

// Create data for admin => middleware.Jwt
// Create data for User
// Login for admin
// Login for User
// Update Username => middleware.Jwt
// Update password => middleware.Jwt
import (
	"BagasA11/GSC-quizHealthEdu-BE/api/controllers"
	"BagasA11/GSC-quizHealthEdu-BE/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(group *gin.RouterGroup) {
	uc := controllers.NewUserController()
	ac := controllers.NewAuthController()
	//User Registration
	group.POST("/user/register", uc.CreateUser)
	//Create another admin
	group.POST("/user/admin/create", middleware.JwtAuth(), uc.CreateAdmin)
	//login page user
	group.POST("/user/login", ac.UserLogin)
	// login page for admin
	group.POST("/user/admin/login", ac.AdmiLogin)
	// update password
	group.POST("/user/password", middleware.JwtAuth(), uc.UpdatePassword)
}
