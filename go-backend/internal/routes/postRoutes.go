package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, pc *controllers.ProdPostController) {

	posts := app.Group("/posts")

	posts.Get("/", pc.GetPosts)

	posts.Get("/:id", pc.GetPost)

	posts.Get("/author/:id", pc.GetPostsByAuthor)

	posts.Post("/", pc.CreatePost)

	posts.Put("/:id", pc.LikePost)

	posts.Delete("/:id", pc.DeletePost)
}
