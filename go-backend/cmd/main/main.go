package main

import (
	"somev2/internal/controllers"
	"somev2/internal/initializers"
	"somev2/internal/models"
	"somev2/internal/repositories"
	"somev2/internal/routes"
	"somev2/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Dependency Injection
	// Post
	postRepo := repositories.NewProdPostRepository(initializers.DB)
	postService := services.NewProdPostService(postRepo)
	postController := controllers.NewProdPostController(postService)

	// User
	userRepo := repositories.NewProdUserRepository(initializers.DB)
	userService := services.NewProdUserService(userRepo)
	userController := controllers.NewProdUserController(userService)

	// Routes
	routes.PostRoutes(app, postController)
	routes.UserRoutes(app, userController)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}
