package main

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"webapi/api"
	"webapi/config"
	"webapi/store"
)

func main() {
	fmt.Println("Building a web API with Go!")

	sqldb := store.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Envs.DBAddress, config.Envs.Port),
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	db, err := sqldb.Init()
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	apiServer := api.NewAPIServer(":8080", db)
	apiServer.Serve()
}
