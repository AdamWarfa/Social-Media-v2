package main

import (
	"somev2/internal/controllers"
	"somev2/internal/initializers"
	"somev2/internal/models"
	"somev2/internal/repositories"
	"somev2/internal/routes"
	"somev2/internal/services"
	"somev2/internal/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Validator
	v := validator.New()

	// Middleware
	app.Use(cors.New())

	// Dependency Injection
	// Post
	postRepo := repositories.NewProdPostRepository(initializers.DB)
	postService := services.NewProdPostService(postRepo)
	postController := controllers.NewProdPostController(postService, v)

	// User
	userRepo := repositories.NewProdUserRepository(initializers.DB)
	userService := services.NewProdUserService(userRepo)
	userController := controllers.NewProdUserController(userService, v)

	// Routes
	routes.PostRoutes(app, postController)
	routes.UserRoutes(app, userController)
	routes.NbaRoutes(app)

	// Start server in a goroutine
	go func() {
		// Start server
		err := app.Listen(":4000")
		if err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	utilities.AwaitSignal(app)

}
