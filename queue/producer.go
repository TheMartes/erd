package queue

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
)

func Populate() {
	nsqconfig := nsq.NewConfig()
	producerURL := config.GetEnvValue(envparams.NSQProducerURL)
	producer, err := nsq.NewProducer(producerURL, nsqconfig)

	if err != nil {
		log.Fatalf("err occured %s", err)
	}

	msgBody := []byte(faker.Word())
	topicName := "hello"

	err = producer.Publish(topicName, msgBody)

	if err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
