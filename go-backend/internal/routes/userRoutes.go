package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, uc *controllers.ProdUserController) {

	users := app.Group("/users")

	users.Get("/", uc.GetUsers)

	users.Get("/:id", uc.GetUser)

	users.Post("/", uc.SaveUser)

	users.Put("/:id", uc.UpdateUser)
}
