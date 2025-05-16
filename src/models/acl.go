package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Acl struct {
		ID        uint           `gorm:"primarykey" json:"id"`
		CreatedAt time.Time      `json:"created"`
		UpdatedAt time.Time      `json:"updated"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
		Acc       int8           `json:"acc" validate:"required"`
		UserID    uint           `json:"userID" validate:"required" gorm:"index:idx_user_topic,unique"`
		Topic     string         `json:"topic" validate:"required" gorm:"index:idx_user_topic,unique"`
		User      User           `gorm:"foreignKey:UserID" json:"user"`
	}
)
