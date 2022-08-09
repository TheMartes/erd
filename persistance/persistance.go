package persistance

import (
	"github.com/elastic/go-elasticsearch/v7"
	elasticserviceprovider "github.com/themartes/erd/persistance/elasticsearch"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	elasticClient *elasticsearch.Client
	mongoClient   *mongo.Client
)

func GetElasticClient() *elasticsearch.Client {
	if elasticClient == nil {
		elasticClient = elasticserviceprovider.Elastic{}.GetClient()
	}

	return elasticClient
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		mongoClient = mongoserviceprovider.Mongo{}.GetClient()
	}

	return mongoClient
}
