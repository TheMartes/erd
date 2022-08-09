package mongoserviceprovider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/themartes/erd/env"
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
	url := env.Params.MongoUrl
	username := env.Params.MongoUsername
	password := env.Params.MongoPassword

	return fmt.Sprintf("mongodb://%s:%s@%s", username, password, url)
}
