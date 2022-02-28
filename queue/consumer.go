package queue

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	return errors.New("")
}

func Init() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "test", config)

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
