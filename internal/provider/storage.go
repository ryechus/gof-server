package provider

type Storageable interface {
	GetString(key string) (string, error)
	SetString(key, value string) error
}

type Storage struct {
	strs map[string]string
}

var _ Storageable = &Storage{}

func NewStorage() *Storage {
	return &Storage{strs: make(map[string]string)}
}

func (s *Storage) SetString(key, value string) error {
	s.strs[key] = value
	return nil
}
func (s *Storage) GetString(key string) (string, error) { return s.strs[key], nil }
