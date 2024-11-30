package routes

import (
	"dating-app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	// Routes for authentication
	auth := r.Group("/auth")
	{
		auth.POST("/signup", authController.Signup)
		auth.POST("/login", authController.Login)
	}

	// Routes for user management
	users := r.Group("/users")
	{
		users.GET("/", userController.GetAllUsers)     // Get all users
		users.GET("/:id", userController.GetUserByID) // Get user by ID
	}

	return r
}
