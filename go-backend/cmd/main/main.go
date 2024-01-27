package main

import (
	"somev2/internal/controllers"
	"somev2/internal/initializers"
	"somev2/internal/models"

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

	app.Use(cors.New())

	app.Get("/posts", controllers.GetPosts)

	app.Get("/posts/:id", controllers.GetPost)

	app.Get("/posts/author/:id", controllers.GetPostsByAuthor)

	app.Get("/users", controllers.GetUsers)

	app.Get("/users/:id", controllers.GetUser)

	app.Post("/posts", controllers.CreatePost)

	app.Post("/users", controllers.SaveUser)

	app.Put("/posts/:id", controllers.LikePost)

	app.Put("/users/:id", controllers.UpdateUser)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}

}
