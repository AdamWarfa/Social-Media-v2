package controllers

import (
	"net/http"
	"somev2/internal/models"
	"somev2/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProdPostController struct {
	service services.PostService
}

func NewProdPostController(service services.PostService) *ProdPostController {
	return &ProdPostController{
		service: service,
	}
}

func (pc *ProdPostController) GetPosts(c *fiber.Ctx) error {

	posts, err := pc.service.GetPosts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

func (pc *ProdPostController) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := pc.service.GetPost(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch post"})
	}

	return c.Status(http.StatusOK).JSON(post)
}

func (pc *ProdPostController) GetPostsByAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	posts, err := pc.service.GetPostsByAuthor(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

func (pc *ProdPostController) CreatePost(c *fiber.Ctx) error {
	var body models.Post

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	newUUID := uuid.New()

	post := models.Post{Id: newUUID.String(), Text: body.Text, Author: body.Author, ImgSrc: body.ImgSrc, Likes: body.Likes, PostDate: body.PostDate}

	post, err := pc.service.CreatePost(post)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.Status(http.StatusOK).JSON(post)
}

func (pc *ProdPostController) LikePost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := pc.service.GetPost(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch post"})
	}
	post, err = pc.service.LikePost(id, &post)

	return c.Status(http.StatusOK).JSON(post)
}

func (pc *ProdPostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	err := pc.service.DeletePost(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.SendStatus(http.StatusOK)
}
