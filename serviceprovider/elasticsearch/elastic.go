package esp

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/erd/config"
)

func GetClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(config.Cfg)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	return es
}

func FindOrCreateIndices(name string) {
	es := GetClient()

	es.Indices.Create(name)
	log.Print("Writing to index: ", name)
}
