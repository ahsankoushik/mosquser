package key_value_controller

import (
	"github.com/ahsankoushik/mosquser/src/config"
	controller "github.com/ahsankoushik/mosquser/src/controllers"
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"github.com/ahsankoushik/mosquser/src/models"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/gofiber/fiber/v2"
)

var DB = config.ConnectDB()

func GetKeys(c *fiber.Ctx) error {
	keysArr := c.Context().QueryArgs().PeekMulti("keys")
	var keys []string
	for _, b := range keysArr {
		keys = append(keys, string(b))
	}
	var data []models.KeyValueStore
	DB.Where("key IN ?", keys).Find(&data)
	return c.JSON(dto_res.Response{
		Status:  fiber.StatusOK,
		Message: "",
		Data:    data,
	})
}

func SetKey(c *fiber.Ctx) error {
	var body models.KeyValueStore
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse data")
	}
	if err := controller.Validate.Struct(&body); err != nil {
		msg, data := utils.FormatValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto_res.Response{
			Status:  fiber.StatusBadRequest,
			Message: msg,
			Data:    data,
		})
	}
	var dbkv models.KeyValueStore
	DB.Where(models.KeyValueStore{Key: body.Key}).First(&dbkv)
	if dbkv.ID > 0 {
		dbkv.Value = body.Value
		DB.Save(&dbkv)
		return c.JSON(dto_res.Response{
			Status:  fiber.StatusOK,
			Message: "",
			Data:    dbkv,
		})
	}
	result := DB.Create(&body)
	if result.RowsAffected > 0 {
		return c.JSON(dto_res.Response{
			Status:  fiber.StatusOK,
			Message: "",
			Data:    body,
		})
	} else {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
}
