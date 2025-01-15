package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Label struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"-"`
	Name      string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Label) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

type LabelScopes struct{}

func (u LabelScopes) SearchScope(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("lower(name) like '%' || lower(?) || '%'", search)
		}

		return db
	}
}

// scopes for preload product
func (u LabelScopes) PreloadTodoList(scopes []func(*gorm.DB) *gorm.DB, column ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("TodoList", func(db *gorm.DB) *gorm.DB {
			// return selected column product if there are filter on parameter
			if len(column) > 0 {
				return db.Scopes(scopes...).Select(column)
			}

			return db.Scopes(scopes...)
		})
	}
}
