package dto_req

import "gorm.io/gorm"

type (
	Paginator struct {
		Page    int
		PerPage int
	}
)

func UserPaginate(p *Paginator) func(db *gorm.DB) *gorm.DB {
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
