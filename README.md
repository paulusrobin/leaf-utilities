# Leaf-Utilities

## Utilities Project
Supporting utilities for [Leaf Framework](https://github.com/paulusrobin/leaf)

## TODO
- [X] [Config](config)
- [X] [Encoding](encoding)
    - [X] [Binary](encoding/binary)
    - [X] [Json](encoding/json)
- [X] [Mandatory](mandatory)
- [X] [Time](time)
- [X] [Logging](logger/logger)
    - [X] [Uber Zap](logger/integrations/zap)
    - [X] [Logrus](logger/integrations/logrus)
- [x] [Message Queue](messageQueue/messageQueue)
    - [x] [Kafka](messageQueue/integrations/kafka)
    - [x] [Google Pub/Sub](messageQueue/integrations/googlePubsub)
- [x] [Database](database)
    - [x] [SQL](database/sql)
        - [x] [MySQL](database/sql/integrations/gorm/mysql)
        - [x] [Postgres](database/sql/integrations/gorm/postgresql)
    - [x] [NoSQL](database/nosql/nosql)
        - [x] [MongoDB](database/nosql/integrations/gomongo)
- [x] [Cache](cache/cache)
    - [x] [Redis](cache/integrations/redis)
    - [x] [Memcache](cache/integrations/memcache)
- [x] [Application Runner](appRunner)
    - [ ] HTTP
    - [ ] gRPC
    - [ ] Messaging Queue
    - [ ] Worker
- [x] [Web Client](webClient/webClient)
    - [x] [Heimdall + Circuit Breaker](webClient/integrations/heimdall)
- [ ] gRPC Client
    - [ ] 
- [x] [Tracer](tracer/tracer)
   - [ ] [Elastic APM](tracer/integrations/elastic)
   - [x] [Newrelic APM](tracer/integrations/newRelic)
   - [x] [Sentry APM](tracer/integrations/sentry)
- [x] [Migration CLI](leafMigration)