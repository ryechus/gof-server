package handlers

type flagVariation struct {
	Name  string `json:"name"    validate:"required"`
	Value any    `json:"value"    validate:"required"`
}

type createFlagPayload struct {
	Key        string          `json:"key"    validate:"required"`
	FlagType   string          `json:"flag_type"    validate:"required"`
	Variations []flagVariation `json:"variations"    validate:"min=2,required"`
}
