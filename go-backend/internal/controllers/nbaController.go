package controllers

import (
	"net/http"
	"somev2/internal/api"

	"github.com/gofiber/fiber/v2"
)

func GetGamesByDate(c *fiber.Ctx) error {
	date := c.Params("date")
	games := api.GetGamesByDate(date)
	return c.Status(http.StatusOK).JSON(games)
}
