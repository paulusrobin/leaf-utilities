# Leaf Logger - Logrus

## logging
Wrapping sirupsen logrus logging library to follow interface leaf framework

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/logger/logger"
    "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
    ...
)
```
To Initialize logger:
```go
var logger = leafLogrus.New() 
```
or pass the logger config
```go
var logger = leafLogrus.New(...leafLogrus.Option) 
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
INFO[0000] [=Trace-ID=] Hello World                      attributes="map[data:test]" mandatory="map[authorization:map[api_key: service_id: token:***] device:map[app_version: brand: device_id: family: model:] device_type:map[] ip_addresses:[] operating_system:map[family: major: minor: name: patch: patch_minor: version:] trace_id:=Trace-ID= user:map[email:email id:1 is_login:true] user_agent:map[family: major: minor: patch: value:]]" message="[=Trace-ID=] Hello World" timestamp="2022-02-23T15:24:08+07:00"
```