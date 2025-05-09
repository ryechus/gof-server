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

type ContextEvaluation struct {
	Kind       string         `json:"kind"`
	Attributes map[string]any `json:"attributes"`
}

type RuleContext struct {
	ContextKind string `json:"kind"`
	Attribute   string `json:"attribute"`
	Values      []any  `json:"values"`
	Operator    string `json:"operator"`
}

type PutRule struct {
	FlagUUID      string        `json:"flag_uuid"`
	UUID          string        `json:"uuid"`
	Name          string        `json:"name" validate:"required"`
	VariationUUID string        `json:"variation_uuid" validate:"required"`
	Priority      int           `json:"priority"`
	RuleContexts  []RuleContext `json:"contexts" validate:"required"`
}

type UpdateFlagVariation struct {
	FlagKeyUUID string `json:"flag_uuid"`
	UUID        string `json:"uuid" validate:"required"`

	FlagVariation
}
