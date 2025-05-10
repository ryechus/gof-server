package repositories

import (
	"fmt"
	"log"
	"time"

	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FlagVariationRepository[T comparable] struct {
	DB *gorm.DB
}

func (fvr *FlagVariationRepository[T]) GetFlagKeyVariationByUUID(variationUUID datatypes.UUID) (database.FlagVariation[T], error) {
	db := fvr.DB
	var flagVariation database.FlagVariation[T]
	scope := database.GetTableName(flagVariation)(db)
	query := fmt.Sprintf("SELECT uuid, flag_key_uuid, name, value FROM %s WHERE uuid = ?",
		scope.Statement.Table)
	result := db.Raw(query, variationUUID.String()).Scan(&flagVariation)
	if result.RowsAffected == 0 {
		return flagVariation, result.Error
	}
	return flagVariation, nil
}

func (fvr *FlagVariationRepository[T]) CreateFlagKeyVariation(newFlag FlagKey, value payloads.FlagVariation, tx *gorm.DB) (database.FlagVariation[T], *gorm.DB) {
	db := tx
	variationUUID := datatypes.NewUUIDv4()
	castedValue, ok := value.Value.(T)
	if !ok {
		log.Panicf("%+v is not of type %s", value.Value, newFlag.FlagType)
	}
	variation := database.FlagVariation[T]{
		UUID:        variationUUID,
		FlagKeyUUID: newFlag.UUID,
		Value:       castedValue,
		Name:        value.Name,
	}
	now := time.Now().UTC()
	scope := database.GetTableName(variation)(db)
	query := fmt.Sprintf("INSERT INTO %s (uuid, flag_key_uuid, value, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", scope.Statement.Table)
	result := db.Raw(query, variation.UUID, variation.FlagKeyUUID, variation.Value, variation.Name, now, now).Scan(&variation)

	log.Printf("created flag key variation %s with value %v for flag key %s\n", variationUUID, value.Value.(T), newFlag.UUID.String())

	return variation, result
}

func (fvr *FlagVariationRepository[T]) GetFlagVariationValue(variationUUID datatypes.UUID) (T, error) {
	db := fvr.DB
	var flagVariation database.FlagVariation[T]
	scope := database.GetTableName(flagVariation)(db)

	var returnVal T
	query := fmt.Sprintf("SELECT uuid, flag_key_uuid, name, value FROM %s WHERE uuid = ?",
		scope.Statement.Table)
	result := db.Raw(query, variationUUID.String()).Scan(&flagVariation)
	if result.RowsAffected != 0 {
		returnVal = flagVariation.Value
		return returnVal, nil
	}
	return returnVal, result.Error
}

func (fvr *FlagVariationRepository[T]) GetFlagVariations(flagKeyUUID datatypes.UUID) ([]database.FlagVariation[T], error) {
	db := fvr.DB
	var flagVariations []database.FlagVariation[T]
	scope := database.GetTableName(database.FlagVariation[T]{})(db)
	query := fmt.Sprintf("SELECT uuid, flag_key_uuid, name, value FROM %s WHERE flag_key_uuid = ?",
		scope.Statement.Table)
	result := db.Raw(query, flagKeyUUID.String()).Scan(&flagVariations)
	if result.RowsAffected == 0 {
		return flagVariations, result.Error
	}
	return flagVariations, nil
}

func (fvr *FlagVariationRepository[T]) UpdateFlagVariation(flagKeyVariation database.FlagVariation[T], tx *gorm.DB) (database.FlagVariation[T], *gorm.DB) {
	db := tx
	now := time.Now().UTC()
	scope := database.GetTableName(database.FlagVariation[T]{})(db)
	query := fmt.Sprintf("UPDATE %s SET name=$1, value=$2, updated_at=$3 WHERE uuid = $4", scope.Statement.Table)
	result := db.Raw(query, flagKeyVariation.Name, flagKeyVariation.Value, now, flagKeyVariation.UUID).Scan(&flagKeyVariation)
	return flagKeyVariation, result
}
