package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	type user struct {
		ID uuid.UUID `gorm:"type:uuid"`
	}

	type label struct {
		ID        uuid.UUID `gorm:"type:uuid"`
		UserID    uuid.UUID `gorm:"type:uuid"`
		User      user      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		Name      string    `gorm:"type:varchar(255)"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	type todo_list_item struct {
		ID         uuid.UUID `gorm:"type:uuid"`
		TodoListID uuid.UUID `gorm:"type:uuid"`
		Name       string    `gorm:"type:varchar(250)"`
		IsChecked  bool
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	type todo_list struct {
		ID            uuid.UUID `gorm:"type:uuid"`
		UserID        uuid.UUID `gorm:"type:uuid"`
		User          user      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		Label         label     `gorm:"foreignKey:LabelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		LabelID       uuid.UUID `gorm:"type:uuid"`
		IsPinned      bool
		Title         string
		TodoListItems []*todo_list_item `gorm:"foreignKey:TodoListID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	newMigration := gormigrate.Migration{
		ID: "0002",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(label{}); err != nil {
				return err
			}

			if err := tx.AutoMigrate(todo_list_item{}); err != nil {
				return err
			}

			if err := tx.AutoMigrate(todo_list{}); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable(label{}); err != nil {
				return err
			}

			if err := tx.Migrator().DropTable(todo_list_item{}); err != nil {
				return err
			}

			return tx.Migrator().DropTable(todo_list{})
		},
	}
	migrations = append(migrations, &newMigration)
}
