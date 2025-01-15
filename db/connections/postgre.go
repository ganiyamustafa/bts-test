package connections

import (
	"fmt"

	"github.com/ganiyamustafa/bts/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Postgre *gorm.DB

func ConnectPostgre() error {
	user := utils.Env("POSTGRES_DB_USER")
	password := utils.Env("POSTGRES_DB_PASS")
	host := utils.Env("POSTGRES_DB_HOST")
	port := utils.Env("POSTGRES_DB_PORT")
	database := utils.Env("POSTGRES_DB_DATABASE")

	// postgres://user:password@host:port/database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	Postgre = db

	return nil
}
