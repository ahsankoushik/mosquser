package users_controller

import (
	"github.com/ahsankoushik/mosquser/src/config"
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	dto_req "github.com/ahsankoushik/mosquser/src/dto/request"
	"github.com/ahsankoushik/mosquser/src/models"
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
