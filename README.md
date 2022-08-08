<p align="center"><img src="https://raw.githubusercontent.com/TheMartes/erd/readme-update/.github/readme/logo.png" width="500"></p>

# ERD - Elasticsearch Replication Daemon
fast, scalable & resilient replication daemon for elasticsearch

## Features
- Unmatched performance **
- Scalable to the skies
- Support for MongoDB by default
- Support for any data, if u can put them into our NSQ

** Initial load can process up to 1 Milion messagges per second, Queue consumption performance depends on your hardware

## Prerequisites
- [NSQ Cluster](https://nsq.io/)
- [Redis](https://redis.com/)
- [ElasticSearch Cluster](https://www.elastic.co/elasticsearch/)

## Local Development
Local Development is powered by docker. Please see `Makefile` to get you up & running quickly. For local development be sure to set `APP_ENV=dev` in `.env`.

### .env
To successfully start daemon you need to set your `.env` variables. All of the required variables could be found in `.env.dist`

### Running the project
After everything is set just run `make up`

## DIY
By default daemon is handling replication from DB to NSQ so consumer can process data to elasticsearch. You can fill up NSQ by yourself if you need neccessary data manipulation on your side. Don't forget to turn off the producers in config afterwards.

## Roadmap
Our roadmap could be found here. [ERD | Height App](https://height.app/FqK7WCMWtj/daemon)

## License
This project has a [MIT license](https://github.com/TheMartes/erd/blob/dev/LICENSE.md).
