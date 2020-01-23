package config

import (
	"github.com/99ddd/god3-core-config/storage"
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
)

// Get gets the config.
func Get() Config {
	return globalCfg
}
