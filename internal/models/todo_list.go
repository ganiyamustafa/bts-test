package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoList struct {
	ID            uuid.UUID       `json:"id"`
	UserID        uuid.UUID       `json:"-"`
	User          User            `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Label         *Label          `json:"label,omitempty" gorm:"foreignKey:LabelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LabelID       *uuid.UUID      `json:"-"`
	IsPinned      bool            `json:"is_pinned"`
	Title         string          `json:"title"`
	TodoListItems []*TodoListItem `json:"todo_list_items" gorm:"foreignKey:TodoListID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

func (u *TodoList) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

type TodoListScopes struct{}

func (u TodoListScopes) SearchScope(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("lower(title) like '%' || lower(?) || '%'", search)
		}

		return db
	}
}

// scopes for preload product
func (u TodoListScopes) PreloadTodoListItem(scopes []func(*gorm.DB) *gorm.DB, column ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("TodoListItems", func(db *gorm.DB) *gorm.DB {
			// return selected column product if there are filter on parameter
			if len(column) > 0 {
				return db.Scopes(scopes...).Select(column)
			}

			return db.Scopes(scopes...)
		})
	}
}

// scopes for preload product
func (u TodoListScopes) PreloadLabel(scopes []func(*gorm.DB) *gorm.DB, column ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Label", func(db *gorm.DB) *gorm.DB {
			// return selected column product if there are filter on parameter
			if len(column) > 0 {
				return db.Scopes(scopes...).Select(column)
			}

			return db.Scopes(scopes...)
		})
	}
}
