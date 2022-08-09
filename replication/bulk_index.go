package replication

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/themartes/erd/env"
	"github.com/themartes/erd/persistance"
)

// ReplicateBulkIndex :)
func ReplicateBulkIndex(replicationData []string) {
	var (
		docs            []*Doc
		countSuccessful uint64
		err             error
	)

	indicesName := env.Params.ReplicationIndex
	client := persistance.GetElasticClient()

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indicesName,
		Client:        client,
		NumWorkers:    1,
		FlushBytes:    int(5e+6),
		FlushInterval: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	for i := 0; i < len(replicationData); i++ {
		docs = append(docs, &Doc{
			ID:    i,
			Title: replicationData[i],
		})
	}

	start := time.Now().UTC()

	for _, a := range docs {
		// Prepare the data payload: encode article to JSON
		data, err := json.Marshal(a)

		if err != nil {
			log.Fatalf("Cannot encode Document %d: %s", a.ID, err)
		}

		// Add an item to the BulkIndexer
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	dur := time.Since(start).Milliseconds()

	log.Println("Replication of", bi.Stats().NumIndexed, "documents Done in", dur, "ms")
}
