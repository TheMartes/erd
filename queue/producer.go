package queue

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/nsqio/go-nsq"
)

func Populate() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("nsqlookupd:4161", config)

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
