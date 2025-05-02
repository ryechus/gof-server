package storage

import "github.com/placer14/gof-server/internal/handlers/payloads"

type Storageable interface {
	EvaluateFlag(key string, payload payloads.GetFlag) (any, error)

	UpdateFlag(payload payloads.UpdateFlag) error

	PutRule(payload payloads.PutRule) error
}
