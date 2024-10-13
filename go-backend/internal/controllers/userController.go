package controllers

import (
	"fmt"
	"net/http"
	"somev2/internal/models"
	"somev2/internal/services"
	"somev2/internal/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// UserController is a contract for the UserController
type UserControllerI interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

// UserController is a struct for the UserController
type UserController struct {
	service  services.UserServiceI
	logger   *zap.Logger
	validate *validator.Validate
	UserControllerI
}

// NewProdUserController is a constructor for the UserController
func NewUserController(service services.UserServiceI, validate *validator.Validate) *UserController {
	return &UserController{
		service:  service,
		logger:   utilities.NewLogger(),
		validate: validate,
	}
}

// GetUsers is a method to get all users
func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetUsers()
	if err != nil {
		uc.logger.Error("Failed to fetch users (controller)", zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.Status(http.StatusOK).JSON(users)
}

// GetUser is a method to get a user by id
func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uc.service.GetUser(id)
	if err != nil {
		uc.logger.Error("Failed to fetch user (controller)", zap.String("id", id), zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		uc.logger.Error("Error parsing register request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := uc.validate.Struct(req); err != nil {
		uc.logger.Error("Error validating register request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := uc.service.RegisterUser(req.Username, req.Email, req.Password, req.Avatar)
	if err != nil {
		uc.logger.Error("Error registering user", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully!",
	})
}

// UpdateUser is a method to update a user in the database
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var body models.User

	if err := c.BodyParser(&body); err != nil {
		uc.logger.Error("Invalid JSON on update user (controller)", zap.String("id", id), zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if err := uc.validate.Struct(body); err != nil {
		uc.logger.Error("Validation error on update user (controller)", zap.String("id", id), zap.Error(err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uc.service.UpdateUser(id, body)
	if err != nil {
		uc.logger.Error("Failed to update user in database (controller)", zap.String("id", id), zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	fmt.Printf("User %s updated in DB", user.Email)

	uc.logger.Info("User updated in database (controller)", zap.String("email", user.Email))
	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		uc.logger.Error("Error parsing login request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := uc.validate.Struct(req); err != nil {
		uc.logger.Error("Error validating login request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Verify the user credentials
	token, username, id, err := uc.service.VerifyLogin(req.Email, req.Password)
	if err != nil {
		// Return an unauthorized status with an error message in JSON
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// On success, return a JSON response with the user info or token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":    token,
		"username": username,
		"id":       id,
	})
}

func (uc *UserController) Logout(c *fiber.Ctx) error {
	return c.SendString("Logged Out")
}
