package repositories

import (
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
	scope := db.Scopes(database.GetTableName(flagVariation))
	result := scope.First(&flagVariation, "uuid = ?", variationUUID.String())
	if result.RowsAffected == 0 {
		return flagVariation, result.Error
	}
	return flagVariation, nil
}

func (fvr *FlagVariationRepository[T]) CreateFlagKeyVariation(newFlag FlagKey, value payloads.FlagVariation) (database.FlagVariation[T], *gorm.DB) {
	db := fvr.DB
	variationUUID := datatypes.NewUUIDv4()
	variation := database.FlagVariation[T]{
		UUID:        variationUUID,
		FlagKeyUUID: newFlag.UUID,
		Value:       value.Value.(T),
		Name:        value.Name,
	}
	result := db.Scopes(database.GetTableName(variation)).Create(variation)
	log.Printf("created flag key variation %s with value %v for flag key %s\n", variationUUID, value.Value.(T), newFlag.UUID.String())

	return variation, result
}

func (fvr *FlagVariationRepository[T]) GetFlagVariationValue(variationUUID datatypes.UUID) (T, error) {
	db := fvr.DB
	var flagVariation database.FlagVariation[T]
	scope := db.Scopes(database.GetTableName(flagVariation))

	var returnVal T
	result := scope.First(&flagVariation, "uuid = ?", variationUUID)
	if result.RowsAffected != 0 {
		returnVal = flagVariation.Value
		return returnVal, nil
	}
	return returnVal, result.Error
}
