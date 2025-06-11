package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ahsankoushik/mosquser/src/config"
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"github.com/ahsankoushik/mosquser/src/models"
	"github.com/ahsankoushik/mosquser/src/routes"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func runServer() {
	app := fiber.New(fiber.Config{
		Prefork: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default to 500 Internal Server Error
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send JSON error response
			return c.Status(code).JSON(dto_res.Response{
				Status:  code,
				Message: err.Error(),
				Data:    fiber.Map{},
			})
		},
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://192.168.0.103:5173",
		AllowCredentials: true,
	}))

	app.Use(logger.New())
	routes.AddRoutes(app)
	host := flag.String("h", "0.0.0.0:6969", "host address" )

	app.Listen(*host)
}

func createSuperUser() {
	var email string
	var password string
	fmt.Println("Enter email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter password: ")
	fmt.Scan(&password)
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatal("User create error", err)
	}
	db := config.ConnectDB()
	user := models.User{
		Email:     email,
		Password:  hashedPassword,
		SuperUser: true,
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Unable to save the user", result.Error)
	}
	log.Println("User created. ID: ", user.ID)
}

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "migrate":
			models.Migrate()
		case "superuser":
			createSuperUser()
		default:
			runServer()
		}

	} else {
		runServer()
	}
}
