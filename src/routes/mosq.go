package routes

import (
	mosq_controller "github.com/ahsankoushik/mosquser/src/controllers/mosq"
	"github.com/gofiber/fiber/v2"
)

func addMosqRoutes(app *fiber.App) {
	mosq := app.Group("/mosq")
	mosq.Post("/auth", mosq_controller.Auth)
	mosq.Post("/superuser", mosq_controller.SuperUser)
	mosq.Post("/acl", mosq_controller.Acl)
}
