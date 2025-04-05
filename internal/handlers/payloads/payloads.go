package payloads

type FlagVariation struct {
	Name  string `json:"name"    validate:"required"`
	Value any    `json:"value"    validate:"required"`
}

type CreateFlagPayload struct {
	Key        string          `json:"key"    validate:"required"`
	FlagType   string          `json:"flag_type"    validate:"required"`
	Variations []FlagVariation `json:"variations"    validate:"min=2,required"`
}
