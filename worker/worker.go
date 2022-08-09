package worker

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/themartes/erd/config"
	"github.com/themartes/erd/env"
	"github.com/themartes/erd/queue"
)

type ReplicationWorker struct {
	DBEngine         string
	SourceDB         string
	SourceCollection string
	ReplicationIndex string
	NSQProducer      *nsq.Producer
	NSQConsumer      *nsq.Consumer
}

func CreateReplicationWorker(engine string, sourcedb string, collection string, replicationIndex string) *ReplicationWorker {
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
	}

	return &worker
}

func (worker ReplicationWorker) StartReplication() {
	go queue.StartProducer(worker.DBEngine, worker.SourceDB, worker.NSQProducer)
	queue.StartConsumer(worker.DBEngine, worker.SourceDB, worker.SourceCollection, worker.NSQConsumer)
}
