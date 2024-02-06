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
	//find user by username
	group.POST("/user/username", middleware.JwtAuth(), uc.FindUsername)
	//find admin by username
	group.POST("/user/admin/username", middleware.JwtAuth(), uc.FindAdminbyUsername)
	//get user by id
	group.GET("/user/all", middleware.JwtAuth(), uc.AllUser)
	//get admin by id
	group.GET("/user/admin/:id", middleware.JwtAuth(), uc.AdminID)
	//get all user data
	group.GET("/user", middleware.JwtAuth(), uc.AllUser)
	//get all type user profile
	group.GET("/user/:id", middleware.JwtAuth(), uc.Me)
	// update password
	group.PUT("/user/password", middleware.JwtAuth(), uc.UpdatePassword)
	//update username ... admin is not allowed to update username
	group.PUT("/user/update-username", middleware.JwtAuth(), uc.UpdateUsername)
	//block user ... admin cannot be blocked
	group.PUT("/user/block/:id", middleware.JwtAuth(), uc.BlockUser)
	//upload avatar image
	group.PUT("/user/avatar", middleware.JwtAuth(), uc.UpdateAvatar)
	//delete user
	group.DELETE("/user/delete", middleware.JwtAuth(), uc.Delete)

}
