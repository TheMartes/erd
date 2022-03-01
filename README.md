# ERD - Elasticsearch Replication Daemon
fast, scalable & resilient replication daemon for elasticsearch

## v0.1

- [ ] Implement Replication Caching system
- [ ] Implement Msg Queue Consumption
- [ ] Split Load into Go routines
- [ ] Create MSG Producer for replication data
- [ ] Split data into 2 Queues, 1 for indexing, 1 for updating
- [ ] Bulk Request after X Documents (Load from config)

## Planned Features (updated - 3.2022)
- Support for Kafka, RabbitMQ
- Replication to different indexes, based on payload data
- Redis Cluster as middleman cache (will be slow tho)
- On demand worker pool extension
- GRPC Support
- /status page
