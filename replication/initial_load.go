package replication

import (
	"runtime"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/erd/helpers"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

var (
	client        *elasticsearch.Client = esp.GetClient()
	indicesName   string
	numberOfCores = runtime.NumCPU()
)

// InitialLoad :)
func InitialLoad(data []string, in string, wg *sync.WaitGroup) {
	chunks := helpers.SplitIntoChunks(data)
	indicesName = in

	for i := 0; i < numberOfCores; i++ {
		chunk := chunks[i]
		go ReplicateBulkIndex(chunk.Data, wg)
	}

	wg.Wait()
}
