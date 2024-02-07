package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func NbaRoutes(app *fiber.App) {

	nba := app.Group("/nba")

	// users.Get("/games", uc.GetUsers)

	nba.Get("/games/:date", controllers.GetGamesByDate)

}
