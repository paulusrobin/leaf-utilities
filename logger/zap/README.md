# Leaf Logger - Zap

## logging
Wrapping uber zap logging library to follow interface leaf framework

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/logger/logger"
    "github.com/paulusrobin/leaf-utilities/logger/zap"
    ...
)
```
To Initialize logger:
```go
var logger = leafZap.New() 
```
or pass the logger config
```go
var logger = leafZap.New(...leafZap.Option) 
```

To logging:
```go
logger.Info(leafLogger.BuildMessage(ctx,
    "Hello World",
    taniLogger.WithAttr("data", "test"),
))
```
output:
```
{"level":"info","ts":"2022-02-23T15:02:33+07:00","caller":"zap/logger.go:19","msg":"[=Trace-ID=] Hello World","message":"[=Trace-ID=] Hello World","attributes":{"data":"test"},"timestamp":"2022-02-23T15:02:33+07:00","mandatory":{"authorization":{"api_key":"","service_id":"","token":"***"},"device":{"app_version":"","brand":"","device_id":"","family":"","model":""},"device_type":{},"ip_addresses":null,"operating_system":{"family":"","major":"","minor":"","name":"","patch":"","patch_minor":"","version":""},"trace_id":"=Trace-ID=","user":{"email":"email","id":1,"is_login":true},"user_agent":{"family":"","major":"","minor":"","patch":"","value":""}}}
```