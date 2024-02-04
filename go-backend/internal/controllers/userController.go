package controllers

import (
	"fmt"
	"net/http"
	"somev2/internal/models"
	"somev2/internal/services"
	"somev2/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// UserController is a contract for the UserController
type UserController interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	SaveUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

// ProdUserController is a struct for the UserController
type ProdUserController struct {
	service services.UserService
	logger  *zap.Logger
}

// NewProdUserController is a constructor for the ProdUserController
func NewProdUserController(service services.UserService) *ProdUserController {
	return &ProdUserController{
		service: service,
		logger:  utilities.NewLogger(),
	}
}

// GetUsers is a method to get all users
func (uc *ProdUserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetUsers()
	if err != nil {
		uc.logger.Error("Failed to fetch users (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.Status(http.StatusOK).JSON(users)
}

// GetUser is a method to get a user by id
func (uc *ProdUserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uc.service.GetUser(id)
	if err != nil {
		uc.logger.Error("Failed to fetch user (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

// SaveUser is a method to save a user in the database
func (uc *ProdUserController) SaveUser(c *fiber.Ctx) error {
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		uc.logger.Error("Invalid JSON on register user (controller)", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user, err := uc.service.SaveUser(body)
	if err != nil {
		uc.logger.Error("Failed to save user in database (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user"})
	}

	fmt.Printf("User %s saved in DB", user.Email)

	uc.logger.Info("User saved in database (controller)", zap.String("email", user.Email))
	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

// UpdateUser is a method to update a user in the database
func (uc *ProdUserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var body models.User

	if err := c.BodyParser(&body); err != nil {
		uc.logger.Error("Invalid JSON on update user (controller)", zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user, err := uc.service.UpdateUser(id, body)
	if err != nil {
		uc.logger.Error("Failed to update user in database (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	fmt.Printf("User %s updated in DB", user.Email)

	uc.logger.Info("User updated in database (controller)", zap.String("email", user.Email))
	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
