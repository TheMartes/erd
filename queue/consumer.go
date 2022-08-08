package queue

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
	"github.com/themartes/erd/replication"
)

var (
	bufferSize = 2000
	msgBuffer  []string
)

// StartConsumer :)
func StartConsumer() {
	nsqconfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(
		config.GetEnvValue(envparams.NSQTopic),
		"consumer",
		nsqconfig,
	)

	if err != nil {
		log.Fatal(err)
	}

	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		if len(m.Body) == 0 {
			return nil
		}

		processMessage(string(m.Body[:]))

		return nil
	}), runtime.NumCPU())

	err = consumer.ConnectToNSQLookupd(config.GetEnvValue(envparams.NSQLookupDaemonURL))

	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	consumer.Stop()
}

func processMessage(payload string) {
	msgBuffer = append(msgBuffer, payload)

	if len(msgBuffer) >= bufferSize {
		data := msgBuffer

		// Clear buffer
		msgBuffer = []string{}

		go replication.ReplicateBulkIndex(data)
	}
}
