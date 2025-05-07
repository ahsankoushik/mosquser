package models

import (
	"time"

	dto_res "github.com/ahsankoushik/mosquser/src/dto"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint           `gorm:"primarykey" json:"id"`
		CreatedAt time.Time      `json:"created"`
		UpdatedAt time.Time      `json:"updated"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
		Email     string         `gorm:"unique,not null,index" json:"email"`
		Password  string         `json:"-"`
		SuperUser bool           `json:"super_user"`
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
