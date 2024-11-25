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

	posts.Get("/:postId/like/count", lc.GetLikeCount)

	protected := app.Group("/api/posts", authMiddeware)

	protected.Post("/:postId/like", lc.LikePost)

	protected.Delete("/:postId/unlike", lc.UnlikePost)

	protected.Get("/:postId/hasliked", lc.HasUserLiked)

}
