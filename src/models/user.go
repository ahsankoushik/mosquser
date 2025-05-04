package models

import (
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
