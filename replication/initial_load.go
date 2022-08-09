package replication

import (
	"runtime"
	"sync"

	"github.com/themartes/erd/helpers"
	"github.com/themartes/erd/persistance"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
)

func InitialLoad(engine string, sourcedb string, collection string, wg *sync.WaitGroup) {
	mongoClient := persistance.GetMongoClient()
	mongoCollection := mongoserviceprovider.GetCollectionFromDB(mongoClient, sourcedb, collection)

	data := mongoserviceprovider.GetDataFromCursor(mongoCollection)

	chunks := helpers.SplitIntoChunks(data)

	for i := 0; i < runtime.NumCPU(); i++ {
		chunk := chunks[i]
		go PutReplicateBulkIndexToWG(chunk.Data, wg)
	}

	wg.Wait()
}
