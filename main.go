package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ahsankoushik/mosquser/src/config"
	"github.com/ahsankoushik/mosquser/src/models"
	"github.com/ahsankoushik/mosquser/src/routes"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func runServer() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	config.ConnectDB()

	app.Use(logger.New())
	app.Static("/static/", "./static")
	routes.AddRoutes(app)

	app.Listen("0.0.0.0:6969")
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
