package storage

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
