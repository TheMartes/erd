package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/themartes/erd/config"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

func main() {
	config.InitEnv()

	// Ultimately this will come from config
	// but for dev purpose we will generate new one
	// each time replication will start
	indicesName := faker.Word()
	esp.FindOrCreateIndices(indicesName)

	var arr []string

	for i := 0; i < 1000000; i++ {
		arr = append(arr, faker.Word())
	}

	fmt.Println("Done.", len(arr))

	//data := replication.GenerateFakeData(500)
	//replication.StartReplicationDaemon(data, indicesName)
}
