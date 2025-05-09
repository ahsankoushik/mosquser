package users_controller

import (
	"github.com/ahsankoushik/mosquser/src/config"
	controller "github.com/ahsankoushik/mosquser/src/controllers"
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	dto_req "github.com/ahsankoushik/mosquser/src/dto/request"
	"github.com/ahsankoushik/mosquser/src/models"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.ConnectDB()

func Get(c *fiber.Ctx) error {
	var body dto_req.Paginator
	c.QueryParser(&body)
	var users []models.User
	DB.Scopes(dto_req.UserPaginate(&body)).Find(&users)
	var count int64
	DB.Model(models.User{}).Count(&count)
	return c.JSON(dto_res.CollectionsRes{
		Status:     fiber.StatusOK,
		Message:    "",
		Page:       body.Page,
		PerPage:    body.PerPage,
		TotalPages: int(count/int64(body.PerPage)) + 1,
		Data:       users,
	})
}

func Create(c *fiber.Ctx) error {
	utils.Logger(string(c.BodyRaw()))
	var body dto_req.CreateUser
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse the data")
	}
	if err := controller.Validate.Struct(&body); err != nil {
		msg, data := utils.FormatValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto_res.Response{
			Status:  fiber.StatusBadRequest,
			Message: msg,
			Data:    data,
		})
	}
	result := DB.Create(&models.User{
		Email:     body.Email,
		Password:  body.Password,
		SuperUser: body.SuperUser,
	})
	if result.Error != nil {
		utils.Logger(result.Error)
	}
	return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
}
