package models

import (
	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Email     string `gorm:"unique,not null,index" validate:"required,email"`
		Password  string `validate:"required,min=8"`
		SuperUser bool
	}
)

func (ar *User) AuthResFromUser(token string) dto_res.AuthRes {
	return dto_res.AuthRes{
		ID:        ar.ID,
		CreatedAt: ar.CreatedAt,
		UpdatedAt: ar.UpdatedAt,
		Email:     ar.Email,
		SuperUser: ar.SuperUser,
		Token:     token,
	}
}
