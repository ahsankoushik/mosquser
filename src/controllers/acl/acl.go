package acl_controller

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
	var acls []models.Acl
	DB.Scopes(dto_req.Paginate(&body)).Preload("User").Find(&acls)
	var count int64
	DB.Model(models.Acl{}).Count(&count)
	return c.JSON(dto_res.CollectionsRes{
		Status:     fiber.StatusOK,
		Message:    "",
		Page:       body.Page,
		PerPage:    body.PerPage,
		TotalPages: int(count/int64(body.PerPage)) + 1,
		Data:       acls,
	})
}

func Create(c *fiber.Ctx) error {
	var body models.Acl
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse data.")
	}
	if err := controller.Validate.Struct(&body); err != nil {
		msg, data := utils.FormatValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto_res.Response{
			Status:  fiber.StatusBadRequest,
			Message: msg,
			Data:    data,
		})
	}
	result := DB.Create(&body)
	if result.RowsAffected > 0 {
		return c.JSON(dto_res.Response{
			Status:  fiber.StatusOK,
			Message: "",
			Data:    body,
		})
	}
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(dto_res.Response{
			Status:  fiber.StatusConflict,
			Message: result.Error.Error(),
			Data:    utils.FormatDBError(result.Error),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(dto_res.Response{
		Status:  fiber.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    fiber.Map{},
	})
}

func Update(c *fiber.Ctx) error {

	return nil
}
