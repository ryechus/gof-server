package repositories

import (
	"github.com/placer14/gof-server/internal/database"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FlagRepository struct {
	DB *gorm.DB
}

var _ Repository = &FlagRepository{}

type FlagKey = database.FlagKey

// func (fr *FlagRepository) Init(db *gorm.DB) error {
// 	fr.DB = db
// 	return nil
// }

func (fr *FlagRepository) GetFlagKey(key string) (FlagKey, *gorm.DB) {
	var flagKey FlagKey
	db := fr.DB

	result := db.First(&flagKey, "key = ?", key)

	return flagKey, result
}

func (fr *FlagRepository) GetFlagKeyByUUID(flagUUID string) (FlagKey, *gorm.DB) {
	db := fr.DB
	var flagKey FlagKey

	result := db.First(&flagKey, "uuid = ?", flagUUID)

	return flagKey, result
}

func (fr *FlagRepository) CreateFlagKey(flagType, key string) (FlagKey, *gorm.DB) {
	db := fr.DB
	flag_key_uuid := datatypes.NewUUIDv4()
	newFlag := database.FlagKey{
		UUID:     flag_key_uuid,
		FlagType: flagType,
		Key:      key,
		Enabled:  false,
	}
	result := db.Create(newFlag)
	return newFlag, result
}

func (fr *FlagRepository) UpdateFlagKey(flagKey *FlagKey) {
	db := fr.DB
	db.Save(flagKey)
}
