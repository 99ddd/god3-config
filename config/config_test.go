package config

import (
	"github.com/99ddd/god3-core-config/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_CreateDatabaseList(t *testing.T) {
	cfg := Config{}

	cfg.Databases = make(storage.DatabaseList)
	cfg.Databases.MakeDefault(storage.Database{
		Driver:       storage.DBDriverMysql,
		Host:         "127.0.0.1",
		Port:         "3306",
		User:         "root",
		Pwd:          "password",
		DatabaseName: "dbName",
		MaxIdleCon:   0,
		MaxOpenCon:   0,
		File:         "",
		Dsn:          "",
	})

	assert.Equal(t, storage.DBDriverMysql, cfg.Databases.GetDefault().Driver)
}
