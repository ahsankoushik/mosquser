package utils

import (
	"log"
	"runtime"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Logger(msg ...any) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("[%s:%d]\n", file, line)
	log.Println(msg)
}

func FormatValidationError(err error) (string, map[string]string) {
	data := make(map[string]string)
	msg := "Validation error on "
	for _, err := range err.(validator.ValidationErrors) {
		msg += err.Field() + " " + err.Tag() + ", "
		data[strings.ToLower(err.Field())] = err.Tag()
	}

	return msg, data
}
