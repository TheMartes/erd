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

	_, err := es.Indices.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Writing to index: ", name)
}
