package dto_req

import "gorm.io/gorm"

type (
	Search struct {
		Username string
		Limit int
	}
	Paginator struct {
		Page    int
		PerPage int
	}
	CreateUser struct {
		Username string `json:"username" validate:"required"`
		Password  string `json:"password" validate:"required,min=8" `
		SuperUser bool   `json:"super_user"`
	}
	UpdateUser struct {
		CreateUser
		ID uint `json:"id" validate:"required"`
	}
)

func Paginate(p *Paginator) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page <= 0 {
			p.Page = 1
		}

		if p.PerPage <= 0 {
			p.PerPage = 30
		}

		offset := (p.Page - 1) * p.PerPage
		return db.Offset(offset).Limit(p.PerPage)
	}
}
