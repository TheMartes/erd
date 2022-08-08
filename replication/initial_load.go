package replication

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/erd/persistance"
)

var (
	client      *elasticsearch.Client = persistance.GetElasticClient()
	indicesName string
)
