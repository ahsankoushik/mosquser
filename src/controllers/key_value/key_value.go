package key_value_controller

import (
	"github.com/ahsankoushik/mosquser/src/config"
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"github.com/ahsankoushik/mosquser/src/models"
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

// func SetKey(c *fiber.Ctx) error {
// 	var
// }
