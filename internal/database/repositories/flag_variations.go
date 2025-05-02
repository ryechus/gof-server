package repositories

import (
	"fmt"
	"log"

	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FlagVariationRepository[T comparable] struct {
	DB *gorm.DB
}

var _ Repository = &FlagVariationRepository[bool]{}

// func (fr *FlagVariationRepository[T]) Init(db *gorm.DB) error {
// 	fr.DB = db
// 	return nil
// }

func (fvr *FlagVariationRepository[T]) GetFlagKeyVariationByUUID(variationUUID datatypes.UUID) (database.FlagVariation[T], error) {
	db := fvr.DB
	var flagVariation database.FlagVariation[T]
	scope := database.GetTableName(flagVariation)(db)
	query := fmt.Sprintf("SELECT uuid, flag_key_uuid, name, value, last_updated FROM %s WHERE uuid = ?",
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
	variation := database.FlagVariation[T]{
		UUID:        variationUUID,
		FlagKeyUUID: newFlag.UUID,
		Value:       value.Value.(T),
		Name:        value.Name,
	}
	scope := database.GetTableName(variation)(db)
	query := fmt.Sprintf("INSERT INTO %s (uuid, flag_key_uuid, value, name) VALUES (?, ?, ?, ?)", scope.Statement.Table)
	result := db.Raw(query, variation.UUID, variation.FlagKeyUUID, variation.Value, variation.Name).Scan(&variation)

	log.Printf("created flag key variation %s with value %v for flag key %s\n", variationUUID, value.Value.(T), newFlag.UUID.String())

	return variation, result
}

func (fvr *FlagVariationRepository[T]) GetFlagVariationValue(variationUUID datatypes.UUID) (T, error) {
	db := fvr.DB
	var flagVariation database.FlagVariation[T]
	scope := database.GetTableName(flagVariation)(db)

	var returnVal T
	query := fmt.Sprintf("SELECT uuid, flag_key_uuid, name, value, last_updated FROM %s WHERE uuid = ?",
		scope.Statement.Table)
	result := db.Raw(query, variationUUID.String()).Scan(&flagVariation)
	if result.RowsAffected != 0 {
		returnVal = flagVariation.Value
		return returnVal, nil
	}
	return returnVal, result.Error
}
