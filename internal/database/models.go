package database

import (
	"time"
)

type FlagKey struct {
	UUID            string `gorm:"primaryKey"`
	Key             string
	ActiveVariation string
	LastUpdated     *time.Time
}

type FlagKeyStringVariations struct {
	UUID        string `gorm:"primaryKey"`
	FlagKeyUUID string
	Value       string
	LastUpdated *time.Time
}

type FlagKeyIntVariations struct {
	UUID        string `gorm:"primaryKey"`
	FlagKeyUUID string
	Value       int64
	LastUpdated *time.Time
}

type FlagKeyFloatVariations struct {
	UUID        string `gorm:"primaryKey"`
	FlagKeyUUID string
	Value       float64
	LastUpdated *time.Time
}

type FlagKeyBoolVariations struct {
	UUID        string `gorm:"primaryKey"`
	FlagKeyUUID string
	Value       bool
	LastUpdated *time.Time
}
