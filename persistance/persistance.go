package persistance

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v9"
	elasticserviceprovider "github.com/themartes/erd/persistance/elasticsearch"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
	redisserviceprovider "github.com/themartes/erd/persistance/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	elasticClient *elasticsearch.Client
	mongoClient   *mongo.Client
	redisClient   *redis.Client
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

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		redisClient = redisserviceprovider.RedisInstance{}.GetClient()
	}

	pong, err := redisClient.Ping(context.TODO()).Result()
	fmt.Println(pong, err)

	return redisClient
}
