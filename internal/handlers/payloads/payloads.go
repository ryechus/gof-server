package payloads

type FlagVariation struct {
	Name  string `json:"name"    validate:"required"`
	Value any    `json:"value"    validate:"required"`
}
