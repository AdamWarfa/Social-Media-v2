package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	app.Get("/posts", controllers.GetPosts)

	app.Get("/posts/:id", controllers.GetPost)

	app.Get("/posts/author/:id", controllers.GetPostsByAuthor)

	app.Post("/posts", controllers.CreatePost)

	app.Put("/posts/:id", controllers.LikePost)
}
