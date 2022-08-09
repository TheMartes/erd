package main

import (
	"log"
	"runtime"

	env "github.com/themartes/erd/env"
	initenv "github.com/themartes/erd/env/init"
	"github.com/themartes/erd/persistance"
	elasticserviceprovider "github.com/themartes/erd/persistance/elasticsearch"
	"github.com/themartes/erd/worker"
)

var (
	numberOfCores = runtime.NumCPU()
)

func main() {
	log.Println("Number of Workers:", numberOfCores)

	if env.Params.AppEnv == "dev" {
		initenv.InitLocal()
	}

	_, err := persistance.GetElasticClient().Indices.Delete([]string{"_all"}) // @Refactor

	if err != nil {
		log.Fatal(err)
	}

	elasticserviceprovider.FindOrCreateIndices(persistance.GetElasticClient())

	worker := worker.CreateReplicationWorker("mongodb", "fluffy", "buffy", "erd", true)
	worker.StartReplication()
}
