package main

import (
	"fmt"
	god3CfgStorage "github.com/99ddd/god3-core-config/storage"
	"os"
)

func main() {
	fmt.Println(os.Args)

	dbList := make(god3CfgStorage.DatabaseList)
	dbList.MakeDefault(god3CfgStorage.Database{
		Driver:       god3CfgStorage.DBDriverMysql,
		Host:         "127.0.0.1",
		Port:         "3306",
		User:         "root",
		Pwd:          "password",
		DatabaseName: "dbName",
	})

	fmt.Printf("%#v\n", dbList.GetDefault())
}
