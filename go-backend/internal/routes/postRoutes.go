package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, pc *controllers.PostController, lc *controllers.LikeController, authMiddeware fiber.Handler) {

	posts := app.Group("/posts")

	posts.Get("/", pc.GetPosts)

	posts.Get("/:id", pc.GetPost)

	posts.Get("/author/:id", pc.GetPostsByAuthor)

	posts.Post("/", pc.CreatePost)

	posts.Put("/:id", pc.LikePost)

	posts.Delete("/:id", pc.DeletePost)

	posts.Get("/:id/like/count", lc.GetLikeCount)

	posts.Get("/:id/hasliked", lc.HasUserLiked)

	protected := app.Group("/api/posts", authMiddeware)

	protected.Post("/:id/like", lc.LikePost)

	protected.Delete("/:id/unlike", lc.UnlikePost)

}
