package elasticserviceprovider

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/env"
)

type Elastic struct{}

func (e Elastic) GetClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(config.ElasticSearch)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	return es
}

func FindOrCreateIndices(es *elasticsearch.Client) {
	name := env.Params.ReplicationIndex

	_, err := es.Indices.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Writing to index: ", name)
}
