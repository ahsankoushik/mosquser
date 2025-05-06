package auth_controller

import (
	"github.com/ahsankoushik/mosquser/src/config"
	controller "github.com/ahsankoushik/mosquser/src/controllers"
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"github.com/ahsankoushik/mosquser/src/models"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.ConnectDB()

func EnsureUserExists(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		// No token or invalid token
		return fiber.NewError(fiber.StatusUnauthorized, "Unautherized")
	}

	claims, err := utils.VerifyToken(authHeader[7:])
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unautherized")
	}

	var user models.User
	DB.First(&user, claims.ID)
	// log.Println("user found", user)
	// if user.Username != UserClaims.Id {

	if user.ID != claims.Id {
		return fiber.NewError(fiber.StatusUnauthorized, "Unautherized")
	}
	c.Locals("user", claims)

	return c.Next()
}

func Login(c *fiber.Ctx) error {
	var userBody models.User
	if err := c.BodyParser(&userBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse data.")
	}
	if err := controller.Validate.Struct(&userBody); err != nil {
		msg, data := utils.FormatValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(
			dto_res.Response{
				Status:  fiber.StatusBadRequest,
				Message: msg,
				Data:    data,
			},
		)
	}
	var userDB models.User
	DB.Where(&models.User{Email: userBody.Email}).First(&userDB)
	if utils.VerifyPassword(userBody.Password, userDB.Password) && userDB.SuperUser {
		token, err := utils.CreateTokenWithTime(userDB.ID, userDB.Email, 30)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Unable to create token.")
		}
		return c.JSON(dto_res.Response{
			Status:  fiber.StatusOK,
			Message: "",
			Data:    userDB.AuthResFromUser(token),
		})
	}
	return fiber.NewError(fiber.StatusUnauthorized, "User or Password does not matched.")
}

func Refreash(c *fiber.Ctx) error {
	claims, _ := c.Locals("user").(*utils.UserClaims)
	var user models.User
	DB.First(&user, claims.ID)
	token, err := utils.CreateTokenWithTime(user.ID, user.Email, 30)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to create token.")
	}
	return c.JSON(dto_res.Response{
		Status:  fiber.StatusOK,
		Message: "",
		Data:    user.AuthResFromUser(token),
	})
}
