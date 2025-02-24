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
