package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/placer14/gof-server/internal/database"
)

type key_db int

const KeyDBStorage key_db = iota

type DBStorage struct {
}

var _ Storageable = &DBStorage{}

func (s *DBStorage) GetBool(key string) (bool, error) { return GetFlag[bool](key) }
func (s *DBStorage) SetBool(key string, value bool) error {
	variations := []bool{
		value, value, value, value,
	}

	CreateFlag(key, variations)
	return nil
}

func (s *DBStorage) GetInt(key string) (int64, error) { return GetFlag[int64](key) }
func (s *DBStorage) SetInt(key string, value int64) error {
	variations := []int64{
		value, value, value, value,
	}

	CreateFlag(key, variations)
	return nil
}

func (s *DBStorage) GetFloat(key string) (float64, error) { return GetFlag[float64](key) }
func (s *DBStorage) SetFloat(key string, value float64) error {
	variations := []float64{
		value, value, value, value,
	}

	CreateFlag(key, variations)
	return nil
}

func (s *DBStorage) GetString(key string) (string, error) { return GetFlag[string](key) }
func (s *DBStorage) SetString(key, value string) error {
	variations := []string{
		value, value, value, value,
	}

	CreateFlag(key, variations)
	return nil
}

func CreateFlag[T comparable](key string, variations []T) {
	fmt.Printf("Setting flag value for key %s", key)
	db := database.GetDB()
	var flagKey database.FlagKey

	result := db.First(&flagKey, "key = ?", key)

	now := time.Now()
	if result.RowsAffected == 0 {
		flag_key_uuid := uuid.New()
		newFlag := database.FlagKey{
			UUID:        flag_key_uuid.String(),
			Key:         key,
			Enabled:     false,
			LastUpdated: &now,
		}
		db.Create(newFlag)
		for _, value := range variations {
			variation := database.FlagVariation[T]{
				UUID:        uuid.NewString(),
				FlagKeyUUID: newFlag.UUID,
				Value:       value,
				LastUpdated: &now,
			}
			db.Scopes(database.GetTableName(variation)).Create(variation)
			newFlag.DefaultVariation = variation.UUID
		}
		db.Save(newFlag)
	}
}

func GetFlag[T comparable](key string) (T, error) {
	db := database.GetDB()
	var flagKey database.FlagKey
	var returnVal T

	result := db.First(&flagKey, "key = ?", key)
	if result.RowsAffected != 0 {
		var flagVariation database.FlagVariation[T]
		scope := db.Scopes(database.GetTableName(flagVariation))
		result = scope.First(&flagVariation, "uuid = ?", flagKey.DefaultVariation)
		if result.RowsAffected != 0 {
			returnVal = flagVariation.Value
			return returnVal, nil
		}
	}

	return returnVal, result.Error
}

func NewDBStorage() *DBStorage {
	return &DBStorage{}
}
