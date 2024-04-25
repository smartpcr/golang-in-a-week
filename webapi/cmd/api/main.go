package main

import (
	"fmt"

	"webapi/pkg/api"
	"webapi/pkg/config"
	"webapi/pkg/inject"
)

func main() {
	fmt.Println("Building a web API with Go!")

	dbCfg, err := config.GetDatabase()
	if err != nil {
		panic(err)
	}
	dbStore, err := inject.Initialize(dbCfg) //store.NewDbStorage(dbCfg)
	if err != nil {
		panic(err)
	}
	err = dbStore.Init()
	if err != nil {
		panic(err)
	}
	defer dbStore.Close()

	apiConfig := config.GetApiConfig()
	apiServer := api.NewAPIServer(fmt.Sprintf(":%d", apiConfig.Port), dbStore)
	apiServer.Serve()
}
