<p align="center"><img src="https://raw.githubusercontent.com/TheMartes/erd/master/.github/readme/logo.png" width="400"></p>

# ERD - Elasticsearch Replication Daemon

fast, scalable & resilient replication daemon for elasticsearch<br />
**Please note this is still Work In Progress Do not use for prod use at all**

## Features

-   Unmatched performance \*\*
-   Scalable to the skies
-   Support for MongoDB by default
-   Support for any data, if u can put them into our NSQ

\*\* Initial load can process up to 1 Milion messagges per second, Queue consumption performance depends on your hardware

## Prerequisites for production use

-   [NSQ Cluster](https://nsq.io/)
-   [Redis](https://redis.com/)
-   [ElasticSearch Cluster](https://www.elastic.co/elasticsearch/)

## Local Development

Local Development is powered by [Docker](https://www.docker.com/). Please see `Makefile` to get you up & running quickly. For local development be sure to set `APP_ENV=dev` in `.env`.

**ENV Types**<br/>
`dist`: .env file with all options you can configure, use for custom configuration<br />
`local`: jump start for local development

> **Tip:** You can just do `cp .env.local .env` to use preconfigured settings

**To learn more** about local environment please refer to [this wiki](https://github.com/TheMartes/erd/wiki/Local-Development)

### Running the project

Please refer to `Makefile`. It has all of the commands available to you to run the project in different modes.

## DIY

By default daemon is handling replication from DB to NSQ so consumer can process data to elasticsearch. You can fill up NSQ by yourself if you need neccessary data manipulation on your side. Don't forget to turn off the producers in config afterwards.

## Roadmap

Our roadmap could be found here. [ERD | Notion](https://www.notion.so/martes/fde8f651f54c47988bfd4e1bed386a06?v=e777c6f286e643808375a10a8135ed1d)

## Contributions

If you feel like you want to help me, please, pick an issue from [notion](https://www.notion.so/martes/fde8f651f54c47988bfd4e1bed386a06?v=e777c6f286e643808375a10a8135ed1d) which is not assigned and is in to-do column and create new issue here on github, tagging me for more information.

## License

This project has a [MIT license](https://github.com/TheMartes/erd/blob/dev/LICENSE.md).
