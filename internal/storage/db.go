package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/placer14/gof-server/internal/database"
)

type key_db int

const KeyDBStorage key_db = iota

type DBStorage struct {
}

var _ Storageable = &DBStorage{}

func (s *DBStorage) GetBool(key string) (bool, error) {
	db := database.GetDB()

	var flagVariations []database.FlagKeyBoolVariations
	var flagKey database.FlagKey

	result := db.First(&flagKey, "key = ?", key)
	if result.RowsAffected == 0 {
		return false, result.Error
	}
	result = db.Find(&flagVariations, "flag_key_uuid = ?", flagKey.UUID)
	if result.RowsAffected == 0 {
		return false, result.Error
	}

	return flagVariations[0].Value, nil
}
func (s *DBStorage) SetBool(key string, value bool) error { return nil }
func (s *DBStorage) SetBoolVariations(key string, values []bool) error {
	db := database.GetDB()
	var flagKey database.FlagKey

	result := db.First(&flagKey, "key = ?", key)
	now := time.Now()
	if result.RowsAffected == 0 {
		flag_key_uuid := uuid.New()
		newFlag := database.FlagKey{
			UUID:        flag_key_uuid.String(),
			Key:         key,
			LastUpdated: &now,
		}
		flagKey = newFlag
		db.Create(newFlag)
	}

	for _, value := range values {
		flagVariation := database.FlagKeyBoolVariations{
			UUID:        uuid.New().String(),
			FlagKeyUUID: flagKey.UUID,
			Value:       value,
			LastUpdated: &now,
		}
		db.Create(flagVariation)
	}
	return nil
}

func (s *DBStorage) GetInt(key string) (int64, error)     { return 0, nil }
func (s *DBStorage) SetInt(key string, value int64) error { return nil }

func (s *DBStorage) GetFloat(key string) (float64, error)     { return 0.0, nil }
func (s *DBStorage) SetFloat(key string, value float64) error { return nil }

func (s *DBStorage) GetString(key string) (string, error) { return "", nil }
func (s *DBStorage) SetString(key, value string) error    { return nil }
func (s *DBStorage) SetStringVariations(key string, values []string) error {
	db := database.GetDB()
	var flagKey database.FlagKey

	result := db.First(&flagKey, "key = ?", key)
	now := time.Now()
	if result.RowsAffected == 0 {
		flag_key_uuid := uuid.New()
		newFlag := database.FlagKey{
			UUID:        flag_key_uuid.String(),
			Key:         key,
			LastUpdated: &now,
		}
		flagKey = newFlag
		db.Create(newFlag)
	}

	for _, value := range values {
		flagVariation := database.FlagKeyStringVariations{
			UUID:        uuid.New().String(),
			FlagKeyUUID: flagKey.UUID,
			Value:       value,
			LastUpdated: &now,
		}
		db.Create(flagVariation)
	}
	return nil
}

func NewDBStorage() *DBStorage {
	return &DBStorage{}
}
