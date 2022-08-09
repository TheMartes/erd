package queue

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/env"
	"github.com/themartes/erd/replication"
)

var (
	bufferSize = 2000
	msgBuffer  []string
)

// StartConsumer :)
func StartConsumer(qd *QueueDaemon) {
	nsqconfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(
		qd.SourceDB,
		qd.SourceCollection,
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

	err = consumer.ConnectToNSQLookupd(env.Params.NSQLookupDaemonURL)

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
