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

```