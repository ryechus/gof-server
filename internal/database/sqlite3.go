package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("foo2.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&FlagKey{})
	db.AutoMigrate(&FlagKeyStringVariations{})
	db.AutoMigrate(&FlagKeyBoolVariations{})
	db.AutoMigrate(&FlagKeyFloatVariations{})
	db.AutoMigrate(&FlagKeyIntVariations{})
	db.AutoMigrate(&TargetingRule{})
	db.AutoMigrate(&TargetingRuleContext{})

	return db
}
