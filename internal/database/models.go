package database

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FlagKey struct {
	UUID                    datatypes.UUID `gorm:"primaryKey"`
	Name                    *string
	Description             *string
	Key                     string `gorm:"index"`
	FlagType                string
	DefaultVariation        datatypes.UUID
	DefaultEnabledVariation datatypes.UUID
	Enabled                 bool
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type FlagVariation[T comparable] struct {
	UUID        datatypes.UUID `gorm:"primaryKey"`
	FlagKeyUUID datatypes.UUID
	Name        string
	Value       T
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FlagKey     FlagKey `gorm:"foreignKey:FlagKeyUUID"`
}

type TargetingRule struct {
	UUID          datatypes.UUID `gorm:"primaryKey"`
	Name          string
	FlagKeyUUID   datatypes.UUID
	VariationUUID datatypes.UUID
	Attributes    datatypes.JSON
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&FlagKey{})
	db.AutoMigrate(&FlagKeyStringVariations{})
	db.AutoMigrate(&FlagKeyBoolVariations{})
	db.AutoMigrate(&FlagKeyFloatVariations{})
	db.AutoMigrate(&FlagKeyIntVariations{})
	db.AutoMigrate(&TargetingRule{})
	db.AutoMigrate(&TargetingRuleContext{})
}
