package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
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
