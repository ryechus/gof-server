package provider

type Storageable interface {
	GetString(key string) (string, error)
	SetString(key, value string) error
	GetBool(key string) (bool, error)
	SetBool(key string, value bool) error
}

type Storage struct {
	stringValues map[string]string
	boolValues   map[string]bool
}

var _ Storageable = &Storage{}

func NewStorage() *Storage {
	return &Storage{
		stringValues: make(map[string]string), 
		boolValues: make(map[string]bool),
	}
}

func (s *Storage) SetString(key, value string) error {
	s.stringValues[key] = value
	return nil
}
func (s *Storage) GetString(key string) (string, error) { return s.stringValues[key], nil }

func (s *Storage) SetBool(key string, value bool) error {
	s.boolValues[key] = value
	return nil
}
func (s *Storage) GetBool(key string) (bool, error) { return s.boolValues[key], nil }
