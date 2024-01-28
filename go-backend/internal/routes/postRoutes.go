package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {

	posts := app.Group("/posts")

	posts.Get("/", controllers.GetPosts)

	posts.Get("/:id", controllers.GetPost)

	posts.Get("/author/:id", controllers.GetPostsByAuthor)

	posts.Post("/", controllers.CreatePost)

	posts.Put("/:id", controllers.LikePost)
}
