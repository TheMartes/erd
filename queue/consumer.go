package queue

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bxcodec/faker/v3"
	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	return nil
}

func Init() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(faker.Word(), faker.Word(), config)

	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&messageHandler{})

	err = consumer.ConnectToNSQLookupd("nsqlookupd:4161")

	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	consumer.Stop()
}
