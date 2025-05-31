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
		Email     string         `gorm:"uniqueIndex;not null" json:"email"`
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

func (u *User) AfterDelete(tx *gorm.DB) (err error){
	tx.Where("user_id = ?",u.ID).Delete(&Acl{})
	return 
}
