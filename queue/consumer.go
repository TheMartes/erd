package queue

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/env"
	"github.com/themartes/erd/replication"
)

var (
	bufferSize = 2000
	msgBuffer  []string
	timeout    = time.Now()
)

// StartConsumer :)
func StartConsumer(engine string, sourcedb string, collection string, consumer *nsq.Consumer) {
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		if len(m.Body) == 0 {
			return nil
		}

		processMessage(string(m.Body[:]))

		return nil
	}), runtime.NumCPU())

	cron := gocron.NewScheduler(time.UTC)
	_, cronErr := cron.Every(5).Seconds().Do(idleQueueCheck)

	if cronErr != nil {
		log.Fatal(cronErr)
	}

	cron.StartAsync()

	err := consumer.ConnectToNSQLookupd(env.Params.NSQLookupDaemonURL)

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

		timeout = time.Now()
	}
}

func idleQueueCheck() {
	queueTimeout, err := strconv.Atoi(env.Params.NSQForceMessageProcessingTimeout)

	if err != nil {
		log.Fatal(err)
	}

	if time.Since(timeout).Seconds() >= float64(queueTimeout) {
		data := msgBuffer

		// Clear buffer
		msgBuffer = []string{}

		go replication.ReplicateBulkIndex(data)

		timeout = time.Now()

		log.Println("Queue population not high enough, forcing message processing...")
	}
}
