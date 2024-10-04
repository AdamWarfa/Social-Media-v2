package main

import (
	"log"
	"os"
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
	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	// if err := initializers.DB.AutoMigrate(&models.Post{}); err != nil {
	// 	panic(err)
	// }

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// Fiber instance
	app := fiber.New()

	// Validator
	v := validator.New()

	// Middleware
	app.Use(cors.New())

	// Dependency Injection
	// Post
	postRepo := repositories.NewPostRepository(initializers.DB)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewProdPostController(postService, v)

	// User
	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, v)

	// Routes
	routes.PostRoutes(app, postController)
	routes.UserRoutes(app, userController, jwtSecret)
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
