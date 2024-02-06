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

	group.POST("/user/username", middleware.JwtAuth(), uc.FindUsername)
	group.GET("/user/admin/:id", middleware.JwtAuth(), uc.AdminID)
	group.GET("/user", middleware.JwtAuth(), uc.AllUser)
	group.GET("/user/:id", middleware.JwtAuth(), uc.GetUserByID)
	// update password
	group.PUT("/user/password", middleware.JwtAuth(), uc.UpdatePassword)
	group.PUT("/user/update-username", middleware.JwtAuth(), uc.UpdateUsername)
	group.DELETE("/user/delete", middleware.JwtAuth(), uc.Delete)
	group.PUT("/user/block/:id", middleware.JwtAuth(), uc.BlockUser)
	group.PUT("/user/avatar", middleware.JwtAuth(), uc.UpdateAvatar)
	//find user by user id
	group.POST("/user/:id", middleware.JwtAuth(), uc.GetUserByID)

}
