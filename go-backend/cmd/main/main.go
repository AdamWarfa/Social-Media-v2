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
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func init() {
	// Load environment variables
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.User{})

	// Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(limiter.New())

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

	// Start server
	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}
