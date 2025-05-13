package config

import (
	"github.com/placer14/gof-server/internal/storage"
)

var FlagStorageIface storage.Storageable = storage.NewDBStorage()

type FlagStorageType = storage.DBStorage

var KeyVariable = storage.KeyDBStorage
