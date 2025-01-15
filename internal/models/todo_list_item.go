package models

import (
	"time"

	"github.com/ganiyamustafa/bts/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoListItem struct {
	ID         uuid.UUID `json:"id"`
	TodoListID uuid.UUID `json:"-"`
	TodoList   TodoList  `json:"todo_list" gorm:"foreignKey:TodoListID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name       string    `json:"name"`
	IsChecked  bool      `json:"is_checked"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u TodoListItem) FromArrayString(datas []string) []*TodoListItem {
	var items = utils.Map[string, *TodoListItem](datas, func(i int, s string) *TodoListItem {
		return &TodoListItem{
			Name: s,
		}
	})

	return items
}

func (u *TodoListItem) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

type TodoListItemScopes struct{}

func (u TodoListItemScopes) SearchScope(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("lower(name) like '%' || lower(?) || '%'", search)
		}

		return db
	}
}

// scopes for preload product
func (u TodoListItemScopes) PreloadTodoList(scopes []func(*gorm.DB) *gorm.DB, column ...string) func(db *gorm.DB) *gorm.DB {
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
