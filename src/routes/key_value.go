package routes

import (
	auth_controller "github.com/ahsankoushik/mosquser/src/controllers/auth"
	key_value_controller "github.com/ahsankoushik/mosquser/src/controllers/key_value"
	"github.com/gofiber/fiber/v2"
)

func AddKeyValueRoutes(app *fiber.App) {
	kv := app.Group("/key_value", auth_controller.EnsureUserExists)
	kv.Get("/", key_value_controller.GetKeys)
	kv.Post("/", key_value_controller.SetKey)
}
