package routes

import (
	"somev2/internal/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func UserRoutes(app *fiber.App, uc controllers.UserControllerI, jwtSecret string) {

	users := app.Group("/users")

	users.Get("/", uc.GetUsers)

	users.Get("/:id", uc.GetUser)

	users.Post("/register", uc.Register)

	users.Post("/login", uc.Login)

	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
	}))

	users.Put("/:id", uc.UpdateUser)
}
