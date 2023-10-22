package database

import (
	"github.com/izakdvlpr/codepix/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
	"log"
	"os"
)

func ConnectDatabase() *gorm.DB {
	var dsn string
	var database *gorm.DB
	var err error

	dsn = os.Getenv("DATABASE_DSN")
	database, err = gorm.Open(os.Getenv("DATABASE_TYPE"), dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)

		panic(err)
	}

	if os.Getenv("DEBUG") == "true" {
		database.LogMode(true)
	}

	if os.Getenv("DATABASE_AUTO_MIGRATE") == "true" {
		database.AutoMigrate(
			&model.Bank{},
			&model.Account{},
			&model.PixKey{},
			&model.Transaction{},
		)
	}

	return database
}
