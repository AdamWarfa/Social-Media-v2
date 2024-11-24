package main

import (
	"log"
	"os"
	"somev2/internal/controllers"
	"somev2/internal/initializers"
	"somev2/internal/middleware"
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

	if err := initializers.DB.AutoMigrate(&models.Post{}); err != nil {
		panic(err)
	}

	initializers.DB.AutoMigrate(&models.Like{})
	initializers.DB.Exec("ALTER TABLE likes ADD CONSTRAINT unique_user_post_like UNIQUE (user_id, post_id);")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	authMiddleware := middleware.AuthMiddleware(jwtSecret)

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
	postController := controllers.NewPostController(postService, v)

	// User
	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, v)

	// Likes
	likeRepo := repositories.NewLikeRepository(initializers.DB)
	likeService := services.NewLikeService(likeRepo)
	likeController := controllers.NewLikeController(likeService)

	// Routes
	routes.PostRoutes(app, postController, likeController, authMiddleware)
	routes.UserRoutes(app, userController, jwtSecret)

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
