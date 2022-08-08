package elasticserviceprovider

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
)

type Elastic struct{}

func (e Elastic) GetClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(config.Cfg)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	return es
}

func FindOrCreateIndices(es *elasticsearch.Client) {
	name := config.GetEnvValue(envparams.ReplicationIndex)

	_, err := es.Indices.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Writing to index: ", name)
}
