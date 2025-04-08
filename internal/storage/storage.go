package storage

import "github.com/placer14/gof-server/internal/handlers/payloads"

type Storageable interface {
	GetString(key string) (string, error)
	SetString(key, value string) error
	CreateStringFlag(key string, flagType string, variations []payloads.FlagVariation) error
	GetBool(key string) (bool, error)
	SetBool(key string, value bool) error
	CreateBoolFlag(key string, flagType string, variations []payloads.FlagVariation) error
	GetFloat(key string) (float64, error)
	SetFloat(key string, value float64) error
	CreateFloatFlag(key string, flagType string, variations []payloads.FlagVariation) error
	GetInt(key string) (int64, error)
	SetInt(key string, value int64) error
	CreateIntFlag(key string, flagType string, variations []payloads.FlagVariation) error

	UpdateFlag(payload payloads.UpdateFlag) error

	PutRule(payload payloads.PutRule) error
}
