package payloads

type FlagVariation struct {
	Name  string `json:"name"    validate:"required"`
	Value any    `json:"value"    validate:"required"`
}

type CreateFlag struct {
	Key        string          `json:"key"    validate:"required"`
	FlagType   string          `json:"flag_type"    validate:"required"`
	Variations []FlagVariation `json:"variations"    validate:"min=2,required"`
}

type UpdateFlag struct {
	Key     string `json:"key" validate:"required"`
	Enabled bool   `json:"enabled"`
}
