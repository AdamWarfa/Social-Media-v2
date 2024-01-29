package controllers

import (
	"net/http"
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

	post, err := services.GetPost(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch post"})
	}
	post, err = services.LikePost(id, &post)

	return c.Status(http.StatusOK).JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.DeletePost(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.SendStatus(http.StatusOK)
}
