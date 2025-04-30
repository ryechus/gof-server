package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=dbpw dbname=gof_server port=5433 sslmode=disable TimeZone=America/Los_Angeles"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	MigrateDB(db)

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&FlagKey{})
	db.AutoMigrate(&FlagKeyStringVariations{})
	db.AutoMigrate(&FlagKeyBoolVariations{})
	db.AutoMigrate(&FlagKeyFloatVariations{})
	db.AutoMigrate(&FlagKeyIntVariations{})
	db.AutoMigrate(&TargetingRule{})
	db.AutoMigrate(&TargetingRuleContext{})
}
