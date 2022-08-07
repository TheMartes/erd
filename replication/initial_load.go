package replication

import (
	"github.com/elastic/go-elasticsearch/v7"
	esp "github.com/themartes/erd/serviceprovider/elasticsearch"
)

var (
	client      *elasticsearch.Client = esp.GetClient()
	indicesName string
)
