package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)

	app.Get("/users/:id", controllers.GetUser)

	app.Post("/users", controllers.SaveUser)

	app.Put("/users/:id", controllers.UpdateUser)
}
