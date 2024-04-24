package main

import (
	"fmt"

	"webapi/pkg/config"
	"webapi/pkg/store"
)

func main() {
	fmt.Println("Building a web API with Go!")

	dbCfg, err := config.GetDatabase()
	if err != nil {
		panic(err)
	}
	dbStore, err := store.NewDbStorage(dbCfg)
	if err != nil {
		panic(err)
	}
	err = dbStore.Init()
	if err != nil {
		panic(err)
	}
	defer dbStore.Close()

	apiConfig := config.GetApiConfig()
	apiServer := NewAPIServer(fmt.Sprintf(":%d", apiConfig.Port), dbStore)
	apiServer.Serve()
}