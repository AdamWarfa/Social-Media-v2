package controllers

import (
	"fmt"
	"net/http"
	"somev2/internal/models"
	"somev2/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	SaveUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type ProdUserController struct {
	service services.UserService
}

func NewProdUserController(service services.UserService) *ProdUserController {
	return &ProdUserController{
		service: service,
	}
}

func (uc *ProdUserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (uc *ProdUserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uc.service.GetUser(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (uc *ProdUserController) SaveUser(c *fiber.Ctx) error {
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user, err := uc.service.SaveUser(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user"})
	}

	fmt.Printf("User %s saved in DB", user.Email)

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

func (uc *ProdUserController) UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user, err := uc.service.UpdateUser(id, body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	fmt.Printf("User %s updated in DB", user.Email)

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
