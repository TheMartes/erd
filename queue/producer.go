package queue

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/env"
	"github.com/themartes/erd/persistance"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
)

// Start producer
func StartProducer(dbengine string, sourcedb string, producer *nsq.Producer) {
	db := env.Params.MongoDB
	collection := env.Params.MongoCollection

	mongoClient := persistance.GetMongoClient()
	mongoCollection := mongoserviceprovider.GetCollectionFromDB(mongoClient, db, collection)

	producerData := mongoserviceprovider.GetDataFromCursor(mongoCollection)

	topicName := fmt.Sprintf("%s.%s", dbengine, sourcedb)

	for _, msg := range producerData {
		msgBody := []byte(msg)
		err := producer.Publish(topicName, msgBody)

		if err != nil {
			log.Fatal(err)
		}
	}
}

// Populate for local development
func Populate(data []string, dbengine string, sourcedb string) {
	nsqconfig := nsq.NewConfig()
	producerURL := env.Params.NSQProducerURL
	producer, err := nsq.NewProducer(producerURL, nsqconfig)

	if err != nil {
		log.Fatalf("err occured %s", err)
	}

	topicName := fmt.Sprintf("%s.%s", dbengine, sourcedb)

	for _, msg := range data {
		msgBody := []byte(msg)
		err = producer.Publish(topicName, msgBody)

		if err != nil {
			log.Fatal(err)
		}
	}

	producer.Stop()
}
