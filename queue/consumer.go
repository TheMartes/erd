package queue

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bxcodec/faker/v3"
	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
)

type messageHandler struct{}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	return nil
}

func Init() {
	nsqconfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(faker.Word(), faker.Word(), nsqconfig)

	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&messageHandler{})

	err = consumer.ConnectToNSQLookupd(config.GetEnvValue(envparams.NSQLookupDaemonURL))

	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	consumer.Stop()
}
