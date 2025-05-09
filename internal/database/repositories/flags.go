package repositories

import (
	"time"

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
	query := `
		SELECT uuid, key, name, description, flag_type,
			default_variation, default_enabled_variation,
			enabled, created_at, updated_at
		FROM flag_keys
		WHERE key = ?
	`
	result := db.Raw(query, key).Scan(&flagKey)

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
	now := time.Now().UTC()
	newFlag := database.FlagKey{
		UUID:     flag_key_uuid,
		FlagType: flagType,
		Key:      key,
		Enabled:  false,
	}
	query := "INSERT INTO flag_keys (uuid, name, flag_type, key, enabled, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result := db.Raw(query, newFlag.UUID, newFlag.Name, newFlag.FlagType, newFlag.Key, newFlag.Enabled, now, now).Scan(&newFlag)
	return newFlag, result
}

func (fr *FlagRepository) UpdateFlagKey(flagKey *FlagKey, tx *gorm.DB) (*FlagKey, *gorm.DB) {
	db := tx
	now := time.Now().UTC()
	query := `UPDATE flag_keys SET name=$1, description=$2 , default_variation=$3, default_enabled_variation=$4, enabled=$5, updated_at=$6 WHERE uuid = $7`
	result := db.Raw(query, flagKey.Name, flagKey.Description, flagKey.DefaultVariation, flagKey.DefaultEnabledVariation, flagKey.Enabled, now, flagKey.UUID).Scan(&flagKey)

	return flagKey, result
}
