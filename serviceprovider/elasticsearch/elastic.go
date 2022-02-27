package esp

import (
	"fmt"
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

func ConfigureDaemon(indicesName string) {
	es := GetClient()

	es.Indices.Create(indicesName)
	fmt.Println("Indices name:", indicesName)
}
