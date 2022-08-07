package queue

import (
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
)

// Populate :)
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
