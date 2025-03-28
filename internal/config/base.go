package config

import "github.com/placer14/gof-server/internal/storage"

var FlagStorageIface storage.Storageable = storage.NewInMemoryStorage()

type FlagStorageType = storage.InMemoryStorage

var KeyVariable = storage.KeyInMemoryStorage
