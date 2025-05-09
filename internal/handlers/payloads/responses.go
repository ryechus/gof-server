package payloads

import "time"

type FlagRepresentation struct {
	FlagUUID                 string             `json:"flag_uuid"`
	Key                      string             `json:"key"`
	Name                     string             `json:"name"`
	Description              string             `json:"description"`
	FlagType                 string             `json:"flag_type"`
	Enabled                  bool               `json:"enabled"`
	DefaultEnabledVariation  any                `json:"default_enabled"`
	DefaultDisabledVariation any                `json:"default_disabled"`
	Rules                    []FlagRuleResponse `json:"rules"`
	CreatedAt                time.Time          `json:"created_at"`
	UpdatedAt                time.Time          `json:"updated_at"`
}

type FlagVariationResponse struct {
	UUID     string `json:"uuid"`
	FlagUUID string `json:"flag_uuid"`
	Name     string `json:"name"    validate:"required"`
	Value    any    `json:"value"    validate:"required"`
}

type FlagRuleResponse struct {
	FlagUUID      string `json:"flag_uuid"`
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	VariationUUID string `json:"variation_uuid"`
	Priority      int    `json:"priority"`
	RuleContexts  any    `json:"contexts"`
}
