package worker

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/env"
	"github.com/themartes/erd/queue"
	"github.com/themartes/erd/replication"
)

type ReplicationWorker struct {
	DBEngine         string
	SourceDB         string
	SourceCollection string
	ReplicationIndex string
	NSQProducer      *nsq.Producer
	NSQConsumer      *nsq.Consumer
	withInitialLoad  bool
}

func CreateReplicationWorker(engine string, sourcedb string, collection string, replicationIndex string, withInitialLoad bool) *ReplicationWorker {
	topicName := fmt.Sprintf("%s.%s", engine, sourcedb)

	producer, err := nsq.NewProducer(
		env.Params.NSQProducerURL,
		config.NSQ,
	)

	if err != nil {
		log.Fatal(err)
	}

	consumer, err := nsq.NewConsumer(
		topicName,
		collection,
		config.NSQ,
	)

	if err != nil {
		log.Fatal(err)
	}

	worker := ReplicationWorker{
		engine,
		sourcedb,
		collection,
		replicationIndex,
		producer,
		consumer,
		withInitialLoad,
	}

	return &worker
}

func (worker ReplicationWorker) StartReplication() {
	if worker.withInitialLoad {
		log.Println("Starting initial load")
		startInitialLoad := time.Now()

		wg := new(sync.WaitGroup)
		wg.Add(runtime.NumCPU())

		replication.InitialLoad(worker.DBEngine, worker.SourceDB, worker.SourceCollection, wg)

		endIntialLoad := time.Since(startInitialLoad).Milliseconds()
		log.Println("Initial load done in", endIntialLoad, "ms")
	}

	producerCron := gocron.NewScheduler(time.UTC)
	interval, ConvErr := strconv.Atoi(env.Params.NSQProductionTimeout)

	if ConvErr != nil {
		log.Fatal(ConvErr)
	}

	_, err := producerCron.Every(interval).Seconds().Do(func() {
		go queue.StartProducer(worker.DBEngine, worker.SourceDB, worker.NSQProducer)
	})

	if err != nil {
		log.Fatal(err)
	}

	producerCron.StartAsync()
	queue.StartConsumer(worker.DBEngine, worker.SourceDB, worker.SourceCollection, worker.NSQConsumer)
}
