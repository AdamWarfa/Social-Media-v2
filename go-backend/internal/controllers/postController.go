package controllers

import (
	"net/http"
	"somev2/internal/initializers"
	"somev2/internal/models"
	"somev2/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPosts(c *fiber.Ctx) error {

	posts, err := services.GetPosts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := services.GetPost(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch post"})
	}

	return c.Status(http.StatusOK).JSON(post)
}

func GetPostsByAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	posts, err := services.GetPostsByAuthor(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	var body models.Post

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	newUUID := uuid.New()

	post := models.Post{Id: newUUID.String(), Text: body.Text, Author: body.Author, ImgSrc: body.ImgSrc, Likes: body.Likes, PostDate: body.PostDate}

	post, err := services.CreatePost(post)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.Status(http.StatusOK).JSON(post)
}

func LikePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post
	if err := initializers.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}

	post.Likes = post.Likes + 1

	result := initializers.DB.Save(&post)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save post"})
	}

	return c.Status(http.StatusOK).JSON(post)
}
