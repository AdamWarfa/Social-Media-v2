package main

import (
	"somev2/internal/initializers"
	"somev2/internal/models"
	"somev2/internal/routes"

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

	// Routes
	routes.PostRoutes(app)
	routes.UserRoutes(app)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}
