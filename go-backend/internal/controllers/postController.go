package controllers

import (
	"fmt"
	"net/http"
	"somev2/initializers"
	"somev2/models"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	fmt.Println(posts)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)

	var post models.Post

	if err := initializers.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}

	return c.Status(http.StatusOK).JSON(post)
}

func GetPostsByAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	var posts []models.Post

	result := initializers.DB.Where("author = ?", id).Find(&posts)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.Status(http.StatusOK).JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	var body struct {
		Text     string `json:"text"`
		Author   string `json:"author"`
		ImgSrc   string `json:"imgSrc"`
		Likes    int    `json:"likes"`
		PostDate string `json:"postDate"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	newUUID := uuid.New()

	post := models.Post{Id: newUUID.String(), Text: body.Text, Author: body.Author, ImgSrc: body.ImgSrc, Likes: body.Likes, PostDate: body.PostDate}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save post"})
	}

	fmt.Printf("Post %v saved in DB", &post.Id)

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
