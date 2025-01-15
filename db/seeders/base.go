package seeders

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type appSeeder struct {
	ID   string
	Seed func(tx *gorm.DB) error
	Wipe func(tx *gorm.DB) error
}

var seeders []*appSeeder

func Wipe(db *gorm.DB, id ...string) {
	db.Session(&gorm.Session{CreateBatchSize: 100, AllowGlobalUpdate: true})

	for _, seeder := range seeders {
		// check if seeder id is empty or there are match seeder id
		if len(id) <= 0 || (len(id) > 0 && (seeder.ID == id[0] || id[0] == "")) {
			if err := seeder.Wipe(db); err != nil {
				log.Error(err)
			}
		}
	}
}

func Seed(db *gorm.DB, id ...string) {
	db.Session(&gorm.Session{CreateBatchSize: 100, AllowGlobalUpdate: true})

	// create transaction
	tx := db.Begin()

	for _, seeder := range seeders {
		// check if seeder id is empty or there are match seeder id
		if len(id) <= 0 || (len(id) > 0 && (seeder.ID == id[0] || id[0] == "")) {
			if err := seeder.Seed(tx); err != nil {
				tx.Rollback()
				log.Error(err)
			}
		}
	}

	// commit transaction
	tx.Commit()
}
