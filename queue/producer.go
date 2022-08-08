package queue

import (
	"context"
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
	"github.com/themartes/erd/persistance"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// Start producer
func StartProducer() {
	db := config.GetEnvValue(envparams.MongoDB)
	collection := config.GetEnvValue(envparams.MongoCollection)

	mongoClient := persistance.GetMongoClient()
	mongoCollection := mongoserviceprovider.GetCollectionFromDB(mongoClient, db, collection)

	mongoData, err := mongoCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{}
	if err = mongoData.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	var producerData []string

	for _, result := range results {
		producerData = append(producerData, result["title"].(string))
	}

	nsqconfig := nsq.NewConfig()
	producerURL := config.GetEnvValue(envparams.NSQProducerURL)
	producer, err := nsq.NewProducer(producerURL, nsqconfig)

	if err != nil {
		log.Fatalf("err occured %s", err)
	}

	topicName := config.GetEnvValue(envparams.NSQTopic)

	for _, msg := range producerData {
		msgBody := []byte(msg)
		err = producer.Publish(topicName, msgBody)

		if err != nil {
			log.Fatal(err)
		}
	}
}

// Populate for local development
func Populate(data []string) {
	nsqconfig := nsq.NewConfig()
	producerURL := config.GetEnvValue(envparams.NSQProducerURL)
	producer, err := nsq.NewProducer(producerURL, nsqconfig)

	if err != nil {
		log.Fatalf("err occured %s", err)
	}

	topicName := config.GetEnvValue(envparams.NSQTopic)

	for _, msg := range data {
		msgBody := []byte(msg)
		err = producer.Publish(topicName, msgBody)

		if err != nil {
			log.Fatal(err)
		}
	}

	producer.Stop()
}
