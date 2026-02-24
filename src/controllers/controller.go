package controller

import (
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"github.com/ahsankoushik/mosquser/src/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()


func BodyParse[T any](body *T, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse data")
	}
	if err := Validate.Struct(body); err != nil{
		msg, data := utils.FormatValidationError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto_res.Response{
			Status: fiber.StatusBadRequest,
			Message: msg,
			Data: data,
		})
	}
	return nil
}
