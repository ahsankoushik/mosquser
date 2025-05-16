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

func GetList(c *fiber.Ctx) error {
	var body dto_req.Paginator
	c.QueryParser(&body)
	var users []models.User
	DB.Scopes(dto_req.Paginate(&body)).Find(&users)
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
	pasword, _ := utils.HashPassword(body.Password)
	var user = models.User{
		Email:     body.Email,
		Password:  pasword,
		SuperUser: body.SuperUser,
	}
	result := DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(dto_res.Response{
			Status:  fiber.StatusConflict,
			Message: result.Error.Error(),
			Data:    utils.FormatDBError(result.Error),
		})
	}
	if result.RowsAffected > 0 {
		return c.JSON(dto_res.Response{
			Status:  fiber.StatusOK,
			Message: "User created",
			Data:    user,
		})
	}
	return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
}

func Search(c *fiber.Ctx) error {
	var body dto_req.Search
	if err := c.QueryParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse data.")
	}
	if body.Limit < 1 {
		body.Limit = 5
	}
	var users []models.User
	DB.Limit(body.Limit).Where("email LIKE ?", "%"+body.Email+"%").Find(&users)
	return c.JSON(dto_res.Response{
		Status:  200,
		Message: "",
		Data:    users,
	})
}
