package payloads

type ContextEvaluation struct {
	Kind       string         `json:"kind"`
	Key        string         `json:"key"`
	Attributes map[string]any `json:"attributes"`
}
