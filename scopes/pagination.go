package scopes

import (
	"gorm.io/gorm"
)

const (
	DEFAULT_PAGE = 1
	DEFAULT_PAGE_SIZE = 10
)

func Pagination(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
