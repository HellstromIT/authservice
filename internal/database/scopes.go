package database

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Limit   int    `form:"PaginationLimit,default=1" json:"PaginationLimit,omitempty"`
	Page    int    `form:"PaginationPage,default=1" json:"PaginationPage,omitempty"`
	Last_ID int64  `form:"PaginationLast,default=0" json:"PaginationLast,omitempty"`
	Order   string `form:"PaginationOrder,default=asc" json:"PaginationOrder,omitempty"`
}

func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page == 0 {
			p.Page = 1
		}

		switch {
		case p.Limit > 100:
			p.Limit = 100
		case p.Limit <= 0:
			p.Limit = 10
		}

		offset := (p.Page - 1) * p.Limit
		return db.Offset(offset).Limit(p.Limit)
	}
}
