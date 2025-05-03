package payloads

type FlagRepresentation struct {
	FlagUUID                 string   `json:"flag_uuid"`
	Key                      string   `json:"key"`
	Name                     string   `json:"name"`
	Description              string   `json:"description"`
	FlagType                 string   `json:"flag_type"`
	Enabled                  bool     `json:"enabled"`
	DefaultEnabledVariation  struct{} `json:"default_enabled"`
	DefaultDisabledVariation struct{} `json:"default_disabled"`
}
