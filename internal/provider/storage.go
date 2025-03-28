package provider

const KeyStorage key = iota

type Storageable interface {
	GetString(key string) (string, error)
	SetString(key, value string) error
	GetBool(key string) (bool, error)
	SetBool(key string, value bool) error
	GetFloat(key string) (float64, error)
	SetFloat(key string, value float64) error
	GetInt(key string) (int64, error)
	SetInt(key string, value int64) error
}

type Storage struct {
	stringValues map[string]string
	boolValues   map[string]bool
	floatValues  map[string]float64
	intValues    map[string]int64
}

var _ Storageable = &Storage{}

func NewStorage() *Storage {
	return &Storage{
		stringValues: make(map[string]string),
		boolValues:   make(map[string]bool),
		floatValues:  make(map[string]float64),
		intValues:    make(map[string]int64),
	}
}

func (s *Storage) SetString(key, value string) error {
	s.stringValues[key] = value
	return nil
}
func (s *Storage) GetString(key string) (string, error) {
	return s.stringValues[key], nil
}

func (s *Storage) SetBool(key string, value bool) error {
	s.boolValues[key] = value
	return nil
}
func (s *Storage) GetBool(key string) (bool, error) {
	return s.boolValues[key], nil
}

func (s *Storage) SetFloat(key string, value float64) error {
	s.floatValues[key] = value
	return nil
}
func (s *Storage) GetFloat(key string) (float64, error) {
	return s.floatValues[key], nil
}

func (s *Storage) SetInt(key string, value int64) error {
	s.intValues[key] = value
	return nil
}
func (s *Storage) GetInt(key string) (int64, error) {
	return s.intValues[key], nil
}
