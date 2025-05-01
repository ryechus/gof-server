package storage

import "github.com/placer14/gof-server/internal/handlers/payloads"

type Storageable interface {
	GetString(key string) (string, error)
	GetBool(key string) (bool, error)
	GetFloat(key string) (float64, error)
	GetInt(key string) (int64, error)

	EvaluateFlag(key string, payload payloads.GetFlag) (any, error)

	UpdateFlag(payload payloads.UpdateFlag) error

	PutRule(payload payloads.PutRule) error
}
