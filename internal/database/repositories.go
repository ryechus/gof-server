package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GetFlag(flagKey string) (*FlagKey, error) {
	db := GetDB()
	var existing_flag FlagKey

	result := db.First(&existing_flag, "key = ?", flagKey)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existing_flag, nil
}

func GetFlags() (*[]FlagKey, error) {
	db := GetDB()
	var existing_flags []FlagKey

	result := db.Find(&existing_flags)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existing_flags, nil
}

func GetVariations(flagKey string) (*[]FlagKeyStringVariations, error) {
	db := GetDB()
	existing_flag, err := GetFlag(flagKey)
	if err != nil {
		return nil, err
	}
	var existing_variations []FlagKeyStringVariations

	result := db.Where("flag_key_uuid = ?", existing_flag.UUID).Find(&existing_variations)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existing_variations, nil
}

func GetOrCreateFlag(flagKey string) *FlagKey {
	db := GetDB()
	var existing_flag FlagKey

	result := db.First(&existing_flag, "key = ?", flagKey)
	now := time.Now()

	if result.RowsAffected == 0 {
		flag_key_uuid := uuid.New()
		variation_one_uuid := uuid.New()
		variation_two_uuid := uuid.New()
		db.Create(FlagKeyStringVariations{UUID: variation_one_uuid.String(), FlagKeyUUID: flag_key_uuid.String(), Value: "false", LastUpdated: &now})
		db.Create(FlagKeyStringVariations{UUID: variation_two_uuid.String(), FlagKeyUUID: flag_key_uuid.String(), Value: "true", LastUpdated: &now})
		existing_flag = FlagKey{UUID: flag_key_uuid.String(), Key: flagKey, ActiveVariation: variation_one_uuid.String(), LastUpdated: &now}
		db.Create(existing_flag)
		fmt.Printf("wrote record for flag key %s", flagKey)
	}

	return &existing_flag
}

func SetCurrentFlagVariation(flagKey string, activeVariationUUID string) (*FlagKey, error) {
	db := GetDB()
	now := time.Now()
	existing_flag := GetOrCreateFlag(flagKey)
	if activeVariationUUID == existing_flag.ActiveVariation {
		return existing_flag, nil
	}

	var existing_variation FlagKeyStringVariations
	result := db.First(&existing_variation, "uuid = ? AND flag_key_uuid = ?", activeVariationUUID, existing_flag.UUID)
	if result.Error != nil {
		return nil, result.Error
	}
	existing_flag.ActiveVariation = activeVariationUUID
	existing_flag.LastUpdated = &now
	db.Save(&existing_flag)

	return existing_flag, nil
}

func GetCurrentFlagVariation(flagKey string) (*FlagKeyStringVariations, error) {
	db := GetDB()

	var existing_flag FlagKey

	result := db.First(&existing_flag, "key = ?", flagKey)
	if result.Error != nil {
		return nil, result.Error
	}
	active_variation_uuid := existing_flag.ActiveVariation

	var active_variation FlagKeyStringVariations
	result = db.First(&active_variation, "uuid = ?", active_variation_uuid)
	if result.Error != nil {
		return nil, result.Error
	}

	return &active_variation, nil
}

func UpdateFlagVariationValue(flagVariationUUID string, flagVariationValue string) (*FlagKeyStringVariations, error) {
	db := GetDB()

	var existing_variation FlagKeyStringVariations

	result := db.First(&existing_variation, "uuid = ?", flagVariationUUID)
	if result.Error != nil {
		return nil, result.Error
	}
	existing_variation.Value = flagVariationValue
	result = db.Save(&existing_variation)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existing_variation, nil
}
