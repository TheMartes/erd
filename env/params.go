package env

type envParams struct {
	AppEnv                           string
	ElasticsearchURL                 string
	ReplicationIndex                 string
	NSQProducerURL                   string
	NSQLookupDaemonURL               string
	NSQTopic                         string
	MongoUrl                         string
	MongoUsername                    string
	MongoPassword                    string
	MongoCollection                  string
	MongoDB                          string
	NSQForceMessageProcessingTimeout string
}

var Params = envParams{
	AppEnv:                           getValue("APP_ENV"),
	ElasticsearchURL:                 getValue("ELASTICSEARCH_URL"),
	ReplicationIndex:                 getValue("REPLICATION_INDEX"),
	NSQProducerURL:                   getValue("NSQ_PRODUCER_URL"),
	NSQLookupDaemonURL:               getValue("NSQ_LOOKUP_DAEMON_URL"),
	NSQTopic:                         getValue("NSQ_TOPIC"),
	MongoUrl:                         getValue("MONGO_URL"),
	MongoUsername:                    getValue("MONGO_USERNAME"),
	MongoPassword:                    getValue("MONGO_PASSWORD"),
	MongoCollection:                  getValue("MONGO_COLLECTION"),
	MongoDB:                          getValue("MONGO_DB"),
	NSQForceMessageProcessingTimeout: getValue("NSQ_FORCE_MESSAGE_PROCESSING_TIMEOUT"),
}
