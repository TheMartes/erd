package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/bxcodec/faker/v3"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/themartes/offer-search/config"
	"github.com/themartes/offer-search/replication"
	"github.com/themartes/offer-search/servicedaemon"
)

func main() {
	config.InitEnv()

	client := servicedaemon.GetElasticClient()
	indicesName := faker.Word()

	// this will create github indices
	servicedaemon.ConfigureDaemon(indicesName)

	var fakeTitles []string

	for i := 0; i < 500; i++ {
		fakeTitles = append(fakeTitles, faker.Word())
	}

	wg := sync.WaitGroup{}

	for index, title := range fakeTitles {
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

	replication.StartReplicationDaemon()
}
