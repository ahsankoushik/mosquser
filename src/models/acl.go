package models

import (
	"gorm.io/gorm"
)

type (
	Acl struct {
		gorm.Model
		Acc    int8
		UserID uint
		User   User `gorm:"foreignKey:UserID"`
		Topic  string
	}
)
