package queue

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func Populate() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("nsqlookupd:4161", config)

	if err != nil {
		log.Fatalf("err occured %s", err)
	}

	msgBody := []byte("hello")
	topicName := "world"

	err = producer.Publish(topicName, msgBody)

	if err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
