package mongoserviceprovider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func (m Mongo) GetClient() *mongo.Client {
	if m.Client != nil {
		return m.Client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(buildMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	m.Client = client

	return m.Client
}

func GetCollectionFromDB(client *mongo.Client, db string, collection string) *mongo.Collection {
	dbInstance := client.Database(db)

	return dbInstance.Collection(collection)
}

func buildMongoURI() string {
	url := config.GetEnvValue(envparams.MongoUrl)
	username := config.GetEnvValue(envparams.MongoUsername)
	password := config.GetEnvValue(envparams.MongoPassword)

	return fmt.Sprintf("mongodb://%s:%s@%s", username, password, url)
}
