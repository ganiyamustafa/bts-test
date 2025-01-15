package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations []*gormigrate.Migration

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

func Rollback(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}
}
