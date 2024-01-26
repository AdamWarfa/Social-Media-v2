package controllers

import (
	"fmt"
	"net/http"
	"somev2/initializers"
	"somev2/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	result := initializers.DB.Find(&users)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.Status(http.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}

	return c.Status(http.StatusOK).JSON(user)
}

func SaveUser(c *fiber.Ctx) error {
	var body struct {
		Id        string `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Avatar    string `json:"avatar"`
		Followers int    `json:"followers"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user := models.User{Id: body.Id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user"})
	}

	fmt.Printf("User %s saved in DB", user.Email)

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var body struct {
		Id        string `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Avatar    string `json:"avatar"`
		Followers int    `json:"followers"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user := models.User{Id: body.Id, Email: body.Email, Username: body.Username, Password: body.Password, Avatar: body.Avatar, Followers: body.Followers}

	result := initializers.DB.Model(&user).Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
