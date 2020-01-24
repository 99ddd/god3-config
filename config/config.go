package config

import (
	"github.com/99ddd/god3-core-config/storage"
	"sync"
	"sync/atomic"
)

// Config type is the global config of god3 application. It will be
// initialized in the engine.
//
// Ref:
// https://github.com/GoAdminGroup/go-admin/blob/master/modules/config/config.go
type Config struct {
	// An map supports multi database connection. The first
	// element of Databases is the default connection. See the
	// file connection.go.
	Databases storage.DatabaseList `json:"database",yaml:"database",ini:"database"`

	// Used to set as the localize language which show in the
	// interface.
	Language string `json:"language",yaml:"language",ini:"language"`
}

var (
	globalCfg Config
	count     uint32
	lock      sync.Mutex
)

// Set sets the config.
func Set(cfg Config) Config {

	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&count) != 0 {
		panic("can not set config twice")
	}
	atomic.StoreUint32(&count, 1)

	cfg.Language = setDefault(cfg.Language, "", "en")

	globalCfg = cfg

	return cfg
}

// Get gets the config.
func Get() Config {
	return globalCfg
}

func setDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}
