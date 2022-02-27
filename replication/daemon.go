package replication

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

var (
	client      *elasticsearch.Client = esp.GetClient()
	indicesName string
)

func StartReplicationDaemon(data []string, in string) {
	indicesName = in
	replicate(data)
}

func replicate(data []string) {
	wg := sync.WaitGroup{}

	for index, title := range data {
		wg.Add(1)

		go func(index int, title string, indicesName string) {
			defer wg.Done()

			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			req := esapi.IndexRequest{
				Index:      indicesName,
				DocumentID: strconv.Itoa(index),
				Body:       strings.NewReader(b.String()),
			}

			res, err := req.Do(context.Background(), client)

			if err != nil {
				log.Fatalf("err %s", err)
			}

			defer res.Body.Close()
		}(index, title, indicesName)
	}

	wg.Wait()
}
