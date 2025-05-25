package routes

import (
	auth_controller "github.com/ahsankoushik/mosquser/src/controllers/auth"
	users_controller "github.com/ahsankoushik/mosquser/src/controllers/users"
	"github.com/gofiber/fiber/v2"
)

func AddUsersRoutes(app *fiber.App) {
	users := app.Group("/users", auth_controller.EnsureUserExists)
	users.Get("/", users_controller.GetList)
	users.Post("/", users_controller.Create)
	users.Put("/", users_controller.Update)
	users.Get("/search", users_controller.Search)
}
