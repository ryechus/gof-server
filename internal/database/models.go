package database

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FlagKey struct {
	UUID                    datatypes.UUID `gorm:"primaryKey"`
	Key                     string
	FlagType                string
	DefaultVariation        datatypes.UUID
	DefaultEnabledVariation datatypes.UUID
	Enabled                 bool
	LastUpdated             *time.Time
}

type FlagVariation[T comparable] struct {
	UUID        datatypes.UUID `gorm:"primaryKey"`
	FlagKeyUUID datatypes.UUID
	Name        string
	Value       T
	LastUpdated *time.Time
}

type TargetingRule struct {
	UUID          datatypes.UUID `gorm:"primaryKey"`
	Name          string
	FlagKeyUUID   datatypes.UUID
	VariationUUID datatypes.UUID
}

type TargetingRuleContext struct {
	UUID              datatypes.UUID `gorm:"primaryKey"`
	TargetingRuleUUID datatypes.UUID
	ContextKind       string
	Attribute         string
	Operator          string
	Value             string
}

type FlagKeyStringVariations FlagVariation[string]

type FlagKeyIntVariations FlagVariation[int64]

type FlagKeyFloatVariations FlagVariation[float64]

type FlagKeyBoolVariations FlagVariation[bool]

func GetFlagKey(key string) (FlagKey, error) {
	db := GetDB()
	var flagKey FlagKey

	result := db.First(&flagKey, "key = ?", key)
	if result.RowsAffected == 0 {
		return flagKey, result.Error
	}

	return flagKey, nil
}

func GetTableName[T comparable](variation FlagVariation[T]) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		switch any(variation).(type) {
		case FlagVariation[string]:
			return tx.Table("flag_key_string_variations")
		case FlagVariation[int64]:
			return tx.Table("flag_key_int_variations")
		case FlagVariation[float64]:
			return tx.Table("flag_key_float_variations")
		case FlagVariation[bool]:
			return tx.Table("flag_key_bool_variations")
		default:
			panic("no matching table for datatype; can't save to database")
		}
	}
}
