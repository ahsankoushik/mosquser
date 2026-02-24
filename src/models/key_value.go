package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	KeyValueStore struct {
		ID        uint           `gorm:"primarykey" json:"id"`
		CreatedAt time.Time      `json:"created"`
		UpdatedAt time.Time      `json:"updated"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-" `
		Key       string         `gorm:"uniqueIndex;not null" json:"key" validate:"required"`
		Value     string         `json:"value" validate:"required"`
	}
)
