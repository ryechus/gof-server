package storage

import "github.com/placer14/gof-server/internal/handlers/payloads"

type Storageable interface {
	// GetFlagRepresentation(key string)
	EvaluateFlag(key string, payload payloads.EvaluateFlag) (any, error)
	UpdateFlag(payload payloads.UpdateFlag) error
	PutRule(payload payloads.PutRule) error
}
