package mongoserviceprovider

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDataFromCursor(mongoCollection *mongo.Collection) []string {
	mongoData, err := mongoCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{}
	if err = mongoData.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	var output []string

	for _, result := range results {
		output = append(output, result["title"].(string))
	}

	return output
}
