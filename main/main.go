package main

import (
	"log"
	"runtime"

	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
	"github.com/themartes/erd/queue"
	"github.com/themartes/erd/replication"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

var (
	numberOfCores = runtime.NumCPU()
)

func main() {
	indicesName := config.GetEnvValue(envparams.ReplicationIndex)

	var indices []string
	indices = append(indices, "_all")

	log.Println("Number of Workers:", numberOfCores)

	_, err := esp.GetClient().Indices.Delete(indices)

	if err != nil {
		log.Fatal(err)
	}

	esp.FindOrCreateIndices(indicesName)

	var arr []string

	if config.GetEnvValue(envparams.AppEnv) == "dev" {
		arr = replication.GenerateFakeData(50000)
		go queue.Populate(arr)
	}

	queue.StartConsumer()
}
