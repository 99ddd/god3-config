package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabase_CreateDatabaseList(t *testing.T) {
	dbList := make(DatabaseList)

	dbMySQL := Database{
		Driver:       DBDriverMysql,
		Host:         "127.0.0.1",
		Port:         "3306",
		User:         "root",
		Pwd:          "password",
		DatabaseName: "dbName",
		MaxIdleCon:   0,
		MaxOpenCon:   0,
		File:         "",
		Dsn:          "",
	}
	dbList.Add("default", dbMySQL)

	dbPostgresql := Database{
		Driver:       DBDriverPostgresql,
		Host:         "127.0.0.1",
		Port:         "5432",
		User:         "postgres",
		Pwd:          "password",
		DatabaseName: "dbName",
		MaxIdleCon:   0,
		MaxOpenCon:   0,
		File:         "",
		Dsn:          "",
	}
	dbList.Add("ORDBMS", dbPostgresql)

	groupDB := dbList.GroupByDriver()

	assert.Equal(t, 2, len(dbList))

	assert.Equal(t, DBDriverMysql, dbList.GetDefault().Driver)
	assert.Equal(t, DBDriverPostgresql, dbList["ORDBMS"].Driver)
	assert.Equal(t, 1, len(groupDB[DBDriverMysql]))

	dbList.MakeDefault(dbPostgresql)
	assert.Equal(t, DBDriverPostgresql, dbList.GetDefault().Driver)
}
