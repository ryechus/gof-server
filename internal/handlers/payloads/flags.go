package payloads

type FlagVariation struct {
	Name  string `json:"name"    validate:"required"`
	Value any    `json:"value"    validate:"required"`
}

type CreateFlag struct {
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Key         string          `json:"key"    validate:"required"`
	FlagType    string          `json:"flag_type"    validate:"required"`
	Variations  []FlagVariation `json:"variations"    validate:"min=2,required"`
}

type UpdateFlag struct {
	Key                     string `json:"key" validate:"required"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	DefaultVariation        string `json:"default_variation"`
	DefaultEnabledVariation string `json:"default_enabled_variation"`
	Enabled                 bool   `json:"enabled"`
}

type EvaluateFlag struct {
	Context ContextEvaluation `json:"context"`
}

type UpdateFlagVariation struct {
	FlagKeyUUID string `json:"flag_uuid"`
	UUID        string `json:"uuid" validate:"required"`

	FlagVariation
}
