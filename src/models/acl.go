package models

import (
	"time"
)


type Acl struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	Acc       int8      `json:"acc" validate:"required"`
	UserID    uint      `json:"userID" validate:"required" gorm:"index:idx_user_topic,unique"`
	Topic     string    `json:"topic" validate:"required" gorm:"index:idx_user_topic,unique"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}

