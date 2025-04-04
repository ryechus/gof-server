package database

import (
	"time"

	"gorm.io/gorm"
)

type FlagKey struct {
	UUID             string `gorm:"primaryKey"`
	Key              string
	DefaultVariation string
	Enabled          bool
	LastUpdated      *time.Time
}

type FlagVariation[T comparable] struct {
	UUID        string `gorm:"primaryKey"`
	FlagKeyUUID string
	Value       T
	LastUpdated *time.Time
}

type FlagKeyStringVariations FlagVariation[string]

type FlagKeyIntVariations FlagVariation[int64]

type FlagKeyFloatVariations FlagVariation[float64]

type FlagKeyBoolVariations FlagVariation[bool]

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
