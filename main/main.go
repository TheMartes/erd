package main

import (
	"log"
	"runtime"

	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
	"github.com/themartes/erd/environments"
	"github.com/themartes/erd/persistance"
	elasticserviceprovider "github.com/themartes/erd/persistance/elasticsearch"
	"github.com/themartes/erd/queue"
)

var (
	numberOfCores = runtime.NumCPU()
)

func main() {
	log.Println("Number of Workers:", numberOfCores)

	if config.GetEnvValue(envparams.AppEnv) == "dev" {
		environments.Local{}.InitLocalEnv()
	}

	_, err := persistance.GetElasticClient().Indices.Delete([]string{"_all"}) // @Refactor

	if err != nil {
		log.Fatal(err)
	}

	elasticserviceprovider.FindOrCreateIndices(persistance.GetElasticClient())

	queue.StartConsumer()
}
