package models

import (
	"github.com/ganiyamustafa/bts/internal/requests"
	"gorm.io/gorm"
)

type UtilScopes struct{}

// scopes for filter pagination
func (u UtilScopes) PaginateScope(req requests.PaginateRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := ((req.Page - 1) * req.Limit)

		if req.Limit > 0 {
			db = db.Limit(req.Limit).Offset(offset)
		}

		return db
	}
}

// scopes for filter pagination
func (u UtilScopes) OrderByScope(req requests.FilterRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.OrderBy != "" && req.Sort != "" {
			db = db.Order(req.OrderBy + " " + req.Sort)
		}

		return db
	}
}
