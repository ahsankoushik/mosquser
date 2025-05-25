package routes

import (
	acl_controller "github.com/ahsankoushik/mosquser/src/controllers/acl"
	auth_controller "github.com/ahsankoushik/mosquser/src/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AddAclRoutes(app *fiber.App) {
	acl := app.Group("/acls", auth_controller.EnsureUserExists)
	acl.Get("/", acl_controller.GetList)
	acl.Post("/", acl_controller.Create)
	acl.Put("/", acl_controller.Update)
}
