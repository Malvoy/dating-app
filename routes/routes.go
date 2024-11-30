package routes

import (
	"dating-app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/signup", authController.Signup)
		auth.POST("/login", authController.Login)
	}

	users := r.Group("/users")
	{
		users.GET("/", userController.GetAllUsers)     
		users.GET("/:id", userController.GetUserByID) 
	}

	return r
}
