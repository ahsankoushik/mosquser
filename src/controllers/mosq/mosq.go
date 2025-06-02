package mosq_controller

import (
	"github.com/ahsankoushik/mosquser/src/config"
	controller "github.com/ahsankoushik/mosquser/src/controllers"
	response "github.com/ahsankoushik/mosquser/src/dto"
	dto_req "github.com/ahsankoushik/mosquser/src/dto/request"
	"github.com/ahsankoushik/mosquser/src/models"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.ConnectDB()

type UserAcl struct {
	Id    int
	Topic string
	Acc   int8
}

func Auth(c *fiber.Ctx) error {
	var user dto_req.MosqAuthBody
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Response{
				Status:  fiber.StatusBadRequest,
				Message: "Unable to parse data.",
				Data:    fiber.Map{},
			},
		)
	}
	if err := controller.Validate.Struct(&user); err != nil {
		msg, data := utils.FormatValidationError(err)
		return c.Status(400).JSON(response.Response{
			Status:  fiber.StatusBadRequest,
			Message: msg + " here use email in username field",
			Data:    data,
		})
	}
	_, err := utils.VerifyToken(user.Password)
	if err == nil {
		return c.SendStatus(fiber.StatusOK)
	}
	var userdb models.User
	DB.Where(&models.User{Email: user.Username}).First(&userdb)
	if utils.VerifyPassword(user.Password, userdb.Password) {
		return c.Status(fiber.StatusOK).JSON(response.Response{
			Status:  fiber.StatusOK,
			Message: "",
			Data:    fiber.Map{},
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
		Status:  fiber.StatusUnauthorized,
		Message: "Email or Password does not matched.",
		Data:    fiber.Map{},
	})
}

func SuperUser(c *fiber.Ctx) error {
	var data dto_req.MosqAuthBody
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Response{
				Status:  fiber.StatusBadRequest,
				Message: "Unable to parse data.",
				Data:    fiber.Map{},
			},
		)
	}
	var user models.User
	DB.Where(&models.User{Email: data.Username}).First(&user)
	if user.SuperUser {
		return c.SendStatus(200)
	}
	return c.SendStatus(400)
}

func Acl(c *fiber.Ctx) error {
	var data dto_req.MosqAuthBody
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Response{
				Status:  fiber.StatusBadRequest,
				Message: "Unable to parse data.",
				Data:    fiber.Map{},
			},
		)
	}
	utils.Logger(data)
	var userAcl UserAcl
	DB.Table("users").Select("users.id, acls.topic, acls.acc").Joins("inner join acls on users.id = acls.user_id").Where("users.email = ? and acls.topic = ?", data.Username,data.Topic).Scan(&userAcl)
	utils.Logger(userAcl)
	if data.Acc == 1 || data.Acc == 4 {
		if userAcl.Acc == 1 || userAcl.Acc == 3 {
			return c.SendStatus(fiber.StatusOK)
		}
	} else if data.Acc == 2 {
		if userAcl.Acc > 1 {
			return c.SendStatus(fiber.StatusOK)
		}
	}

	return c.SendStatus(400)
}
