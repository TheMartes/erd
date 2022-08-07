package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/replication"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

var (
	numberOfCores = runtime.NumCPU()
)

func main() {
	config.InitEnv()

	// Ultimately this will come from config
	// but for dev purpose we will generate new one
	// each time replication will start
	indicesName := faker.Word()

	var indices []string
	indices = append(indices, "_all")

	wg := new(sync.WaitGroup)
	wg.Add(numberOfCores)

	log.Println("Number of Workers:", numberOfCores)

	_, err := esp.GetClient().Indices.Delete(indices)

	if err != nil {
		log.Fatal(err)
	}

	esp.FindOrCreateIndices(indicesName)

	arr := replication.GenerateFakeData(100000)

	start := time.Now().UTC()

	replication.InitialLoad(arr, indicesName, wg)

	dur := time.Since(start).Milliseconds()

	fmt.Println("Cycle done in", dur, "ms")
}
