# Leaf-Utilities

## Utilities Project
Supporting utilities for [Leaf Framework](https://github.com/paulusrobin/leaf)

## TODO
- [X] [Config](https://github.com/paulusrobin/leaf-utilities/tree/main/config)
- [X] Encoding
    - [X] [Binary](https://github.com/paulusrobin/leaf-utilities/tree/main/encoding/binary)
    - [X] [Json](https://github.com/paulusrobin/leaf-utilities/tree/main/encoding/json)
- [X] [Mandatory](https://github.com/paulusrobin/leaf-utilities/tree/main/mandatory)
- [X] [Time](https://github.com/paulusrobin/leaf-utilities/tree/main/time)
- [X] [Logging](https://github.com/paulusrobin/leaf-utilities/tree/logger/logger)
    - [X] [Uber Zap](https://github.com/paulusrobin/leaf-utilities/tree/logger/zap)
    - [X] [Logrus](https://github.com/paulusrobin/leaf-utilities/tree/logger/logrus)
- [x] [Message Queue](messageQueue/messageQueue)
    - [x] [Kafka](messageQueue/integrations/kafka)
    - [ ] Google Pub/Sub
- [x] [Database](database)
    - [x] [SQL](database/sql)
        - [x] [MySQL](database/sql/integrations/gorm/mysql)
        - [x] [Postgres](database/sql/integrations/gorm/postgresql)
    - [x] [NoSQL](database/nosql/nosql)
        - [x] [MongoDB](database/nosql/integrations/gomongo)
- [x] [Cache](cache/cache)
    - [x] [Redis](cache/integrations/redis)
    - [x] [Memcache](cache/integrations/memcache)
- [ ] Application Runner
    - [ ] HTTP
    - [ ] gRPC
    - [ ] Messaging Queue
    - [ ] Worker
- [x] [Web Client](webClient/webClient)
    - [x] [Heimdall + Circuit Breaker](webClient/integrations/heimdall)
- [ ] gRPC Client
    - [ ] 
 - [ ] Tracer
    - [ ] Elastic APM
    - [ ] Newrelic APM
    - [ ] Sentry APM