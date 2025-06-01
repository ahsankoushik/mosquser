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
	if err := controller.BodyParse(&body,c); err != nil{
		return err
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
	var body dto_req.CreateUser
	if err := controller.BodyParse(&body,c); err != nil{
		return err
	}
	utils.Logger(body)
	return c.SendStatus(200)
}

func Delete(c *fiber.Ctx) error {
	var body dto_req.Delete
	if err := controller.BodyParse(&body,c); err != nil{
		return err
	}
	result := DB.Delete(&models.Acl{ID: body.ID})
	if result.Error != nil{
		data := utils.FormatDBError(result.Error)
		return c.Status(fiber.StatusConflict).JSON(dto_res.Response{
			Status:fiber.StatusConflict,
			Message: "",
			Data: data,
		})
	}else{
		return c.JSON(dto_res.Response{
			Status: fiber.StatusOK,
			Message: "Deleted",
			Data: fiber.Map{},
		})
	}
}
