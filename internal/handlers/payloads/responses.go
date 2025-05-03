package payloads

import "time"

type FlagRepresentation struct {
	FlagUUID                 string    `json:"flag_uuid"`
	Key                      string    `json:"key"`
	Name                     string    `json:"name"`
	Description              string    `json:"description"`
	FlagType                 string    `json:"flag_type"`
	Enabled                  bool      `json:"enabled"`
	DefaultEnabledVariation  any       `json:"default_enabled"`
	DefaultDisabledVariation any       `json:"default_disabled"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}
