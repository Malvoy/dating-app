package main

import (
	"dating-app/controllers"
	"dating-app/models"
	"dating-app/repositories"
	"dating-app/routes"
	"dating-app/services"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("dating-app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(*authService)
	userController := controllers.NewUserController(*userService)

	r := routes.SetupRouter(authController, userController)
	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
