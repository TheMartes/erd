package servicedaemon

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/offer-search/config"
)

func GetElasticClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(config.Cfg)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	return es
}

func ConfigureDaemon(indicesName string) {
	es := GetElasticClient()

	es.Indices.Create(indicesName)
	fmt.Println("Indices name:", indicesName)
}
