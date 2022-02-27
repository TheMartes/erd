package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/themartes/offer-search/config"
	"github.com/themartes/offer-search/replication"
	"github.com/themartes/offer-search/servicedaemon"
)

func main() {
	config.InitEnv()

	// Ultimately this will come from config
	// but for dev purpose we will generate new one
	// each time replication will start
	indicesName := faker.Word()
	servicedaemon.ConfigureDaemon(indicesName)

	data := replication.GenerateFakeData(500)
	replication.StartReplicationDaemon(data, indicesName)
}
