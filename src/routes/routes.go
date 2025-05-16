package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	app.Static("/", "./frontend/dist")
	app.Get("/_", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/dist/index.html")
	})
	addMosqRoutes(app)
	AddAuthRoutes(app)
	AddUsersRoutes(app)
	AddAclRoutes(app)
}
