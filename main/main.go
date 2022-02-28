package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/replication"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

func main() {
	config.InitEnv()

	// Ultimately this will come from config
	// but for dev purpose we will generate new one
	// each time replication will start
	indicesName := faker.Word()
	esp.FindOrCreateIndices(indicesName)

	data := replication.GenerateFakeData(500)
	replication.StartReplicationDaemon(data, indicesName)
}
