package routes

import (
	auth_controller "github.com/ahsankoushik/mosquser/src/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AddAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", auth_controller.Login)
	auth.Use("/", auth_controller.EnsureUserExists)
	auth.Get("/refreash", auth_controller.Refreash)
}
