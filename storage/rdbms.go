package storage

// Database is a type of database connection config.
//
// Because a little difference of different database driver.
// The Config has multiple options but may not be used.
// Such as the sqlite driver only use the File option which
// can be ignored when the driver is mysql.
//
// If the Dsn is configured, when driver is mysql/postgresql/
// mssql, the other configurations will be ignored, except for
// MaxIdleCon and MaxOpenCon.
//
// Ref:
// https://github.com/GoAdminGroup/go-admin/blob/master/modules/config/config.go
type Database struct {
	Driver       string `json:"driver",yaml:"driver",ini:"driver"`
	Host         string `json:"host",yaml:"host",ini:"host"`
	Port         string `json:"port",yaml:"port",ini:"port"`
	User         string `json:"user",yaml:"user",ini:"user"`
	Pwd          string `json:"pwd",yaml:"pwd",ini:"pwd"`
	DatabaseName string `json:"database_name",yaml:"database_name",ini:"database_name"`
	MaxIdleCon   int    `json:"max_idle_con",yaml:"max_idle_con",ini:"max_idle_con"`
	MaxOpenCon   int    `json:"max_open_con",yaml:"max_open_con",ini:"max_open_con"`
	File         string `json:"file",yaml:"file",ini:"file"`
	Dsn          string `json:"dsn",yaml:"dsn",ini:"dsn"`
}

const (
	// DBDriverMysql is a const value of mysql driver.
	DBDriverMysql = "mysql"
	// DBDriverSqlite is a const value of sqlite driver.
	DBDriverSqlite = "sqlite"
	// DBDriverPostgresql is a const value of postgresql driver.
	DBDriverPostgresql = "postgresql"
	// DBDriverMssql is a const value of mssql driver.
	DBDriverMssql = "mssql"
)

// DatabaseList is a map of Database.
type DatabaseList map[string]Database

// GetDefault get the default Database.
func (d DatabaseList) GetDefault() Database {
	return d["default"]
}

// Add add a Database to the DatabaseList.
func (d DatabaseList) Add(key string, db Database) {
	d[key] = db
}

// MakeDefault a default database config
func (d DatabaseList) MakeDefault(db Database) {
	d.Add("default", db)
}

// GroupByDriver group the Databases with the drivers.
func (d DatabaseList) GroupByDriver() map[string]DatabaseList {
	drivers := make(map[string]DatabaseList)
	for key, item := range d {
		if driverList, ok := drivers[item.Driver]; ok {
			driverList.Add(key, item)
		} else {
			drivers[item.Driver] = make(DatabaseList)
			drivers[item.Driver].Add(key, item)
		}
	}
	return drivers
}
