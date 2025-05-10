package payloads

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

type RolloutVariants struct {
	Percentage    int    `json:"percentage" validate:"required"`
	VariationUUID string `json:"variation_uuid" validate:"required"`
}

type RolloutConfig struct {
	Kind      string            `json:"kind" validate:"required"`
	Attribute string            `json:"attribute" validate:"required"`
	Variants  []RolloutVariants `json:"variants" validate:"required"`
}

type PutRolloutRule struct {
	UUID          string        `json:"uuid"`
	FlagUUID      string        `json:"flag_uuid" validate:"required"`
	RolloutType   string        `json:"rollout_type"`
	ContextConfig RolloutConfig `json:"context_config" validate:"required"`
}
