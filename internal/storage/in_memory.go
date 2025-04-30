package storage

import "github.com/placer14/gof-server/internal/handlers/payloads"

type key_in_memory int

const KeyInMemoryStorage key_in_memory = iota

type InMemoryStorage struct {
	stringValues map[string]string
	boolValues   map[string]bool
	floatValues  map[string]float64
	intValues    map[string]int64
}

var _ Storageable = &InMemoryStorage{}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		stringValues: make(map[string]string),
		boolValues:   make(map[string]bool),
		floatValues:  make(map[string]float64),
		intValues:    make(map[string]int64),
	}
}

func (s *InMemoryStorage) SetString(key, value string) error {
	s.stringValues[key] = value
	return nil
}
func (s *InMemoryStorage) CreateStringFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return nil
}
func (s *InMemoryStorage) GetString(key string) (string, error) {
	return s.stringValues[key], nil
}

func (s *InMemoryStorage) SetBool(key string, value bool) error {
	s.boolValues[key] = value
	return nil
}
func (s *InMemoryStorage) CreateBoolFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return nil
}
func (s *InMemoryStorage) GetBool(key string) (bool, error) {
	return s.boolValues[key], nil
}

func (s *InMemoryStorage) SetFloat(key string, value float64) error {
	s.floatValues[key] = value
	return nil
}
func (s *InMemoryStorage) CreateFloatFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return nil
}
func (s *InMemoryStorage) GetFloat(key string) (float64, error) {
	return s.floatValues[key], nil
}

func (s *InMemoryStorage) SetInt(key string, value int64) error {
	s.intValues[key] = value
	return nil
}
func (s *InMemoryStorage) CreateIntFlag(key string, flagType string, variations []payloads.FlagVariation) error {
	return nil
}
func (s *InMemoryStorage) GetInt(key string) (int64, error) {
	return s.intValues[key], nil
}

func (s *InMemoryStorage) UpdateFlag(payload payloads.UpdateFlag) error {
	return nil
}

func (s *InMemoryStorage) PutRule(payload payloads.PutRule) error {
	return nil
}

func (s *InMemoryStorage) EvaluateFlag(key string, payload payloads.GetFlag) (any, error) {
	return nil, nil
}
