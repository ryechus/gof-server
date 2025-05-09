package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *gorm.DB {
	dsn := os.Getenv("GOF_SERVER_POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=dbpw dbname=gof_server port=5433 sslmode=disable TimeZone=America/Los_Angeles"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(15)
	if err != nil {
		log.Fatal(err)
	}

	MigrateDB(db)

	return db
}
