package controllers

import (
	"net/http"
	"somev2/internal/models"
	"somev2/internal/services"
	"somev2/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// PostController is a contract for the PostController
type PostController interface {
	GetPosts(c *fiber.Ctx) error
	GetPost(c *fiber.Ctx) error
	GetPostsByAuthor(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	LikePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
}

// ProdPostController is a struct for the PostController
type ProdPostController struct {
	service services.PostService
	logger  *zap.Logger
}

// NewProdPostController is a constructor for the ProdPostController
func NewProdPostController(service services.PostService) *ProdPostController {
	return &ProdPostController{
		service: service,
		logger:  utilities.NewLogger(),
	}
}

// GetPosts is a method to get all posts
func (pc *ProdPostController) GetPosts(c *fiber.Ctx) error {
	posts, err := pc.service.GetPosts()
	if err != nil {
		pc.logger.Error("Failed to fetch posts (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

// GetPost is a method to get a post by id
func (pc *ProdPostController) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := pc.service.GetPost(id)
	if err != nil {
		pc.logger.Error("Failed to fetch post (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch post"})
	}

	return c.Status(http.StatusOK).JSON(post)
}

// GetPostsByAuthor is a method to get all posts by author
func (pc *ProdPostController) GetPostsByAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	posts, err := pc.service.GetPostsByAuthor(id)

	if err != nil {
		pc.logger.Error("Failed to fetch posts (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

// CreatePost is a method to create a post
func (pc *ProdPostController) CreatePost(c *fiber.Ctx) error {
	var body models.Post

	if err := c.BodyParser(&body); err != nil {
		pc.logger.Error("Invalid JSON", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	newUUID := uuid.New()

	post := models.Post{Id: newUUID.String(), Text: body.Text, Author: body.Author, ImgSrc: body.ImgSrc, Likes: body.Likes, PostDate: body.PostDate}

	post, err := pc.service.CreatePost(post)

	if err != nil {
		pc.logger.Error("Failed to create post (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	pc.logger.Info("Post created (controller)", zap.String("id", post.Id))
	return c.Status(http.StatusOK).JSON(post)
}

// LikePost is a method to like a post
func (pc *ProdPostController) LikePost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := pc.service.GetPost(id)
	if err != nil {
		pc.logger.Error("Failed to fetch liked post (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch post"})
	}

	post, err = pc.service.LikePost(id, &post)
	if err != nil {
		pc.logger.Error("Failed to like post (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to like post"})
	}

	pc.logger.Info("Post liked (controller)", zap.String("id", post.Id))
	return c.Status(http.StatusOK).JSON(post)
}

// DeletePost is a method to delete a post
func (pc *ProdPostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	err := pc.service.DeletePost(id)
	if err != nil {
		pc.logger.Error("Failed to delete post (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	pc.logger.Info("Post deleted (controller)", zap.String("id", id))
	return c.SendStatus(http.StatusOK)
}
