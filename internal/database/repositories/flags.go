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

	// result := db.First(&flagKey, "key = ?", key)
	result := db.Raw("SELECT uuid, key, flag_type, default_variation, default_enabled_variation, enabled FROM flag_keys WHERE key = ?", key).Scan(&flagKey)

	return flagKey, result
}

func (fr *FlagRepository) GetFlagKeyByUUID(flagUUID string) (FlagKey, *gorm.DB) {
	db := fr.DB
	var flagKey FlagKey

	result := db.First(&flagKey, "uuid = ?", flagUUID)

	return flagKey, result
}

func (fr *FlagRepository) CreateFlagKey(flagType, key string, tx *gorm.DB) (FlagKey, *gorm.DB) {
	db := tx
	flag_key_uuid := datatypes.NewUUIDv4()
	newFlag := database.FlagKey{
		UUID:     flag_key_uuid,
		FlagType: flagType,
		Key:      key,
		Enabled:  false,
	}
	query := "INSERT INTO flag_keys (uuid, flag_type, key, enabled) VALUES (?, ?, ?, ?)"
	result := db.Raw(query, newFlag.UUID, newFlag.FlagType, newFlag.Key, newFlag.Enabled).Scan(&newFlag)
	return newFlag, result
}

func (fr *FlagRepository) UpdateFlagKey(flagKey *FlagKey, tx *gorm.DB) {
	db := tx
	db.Save(flagKey)
}
