package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	users := app.Group("/users")

	users.Get("/", controllers.GetUsers)

	users.Get("/:id", controllers.GetUser)

	users.Post("/", controllers.SaveUser)

	users.Put("/:id", controllers.UpdateUser)
}
