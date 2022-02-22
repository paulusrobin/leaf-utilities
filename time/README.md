# Time

## Time Helper for Go
Time helper for golang package, mocking time without using monkeypatch :see_no_evil:.

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/time"
    ...
)
```

To get current time:
```go
var now = leafTime.Now()
```

To mock / freeze current time and reset the mock:
```go
var desiredTime = leafTime.Now()
leafTime.Mock(desiredTime)
defer leafTime.ResetMock()
```

To convert time to UTC:
```go
var utcTime = leafTime.ToUTCTime(leafTime.Now())
```

To convert time to Server Time:
```go
var serverTime = leafTime.ToServerTime(leafTime.Now())
```

To convert time using Common Location (List [Common Location](https://github.com/paulusrobin/leaf-utilities/time/location.go))
```go
wibTime, err := leafTime.ToClientTimeByLocation(leafTime.Now(), leafTime.WIB)
witaTime, err := leafTime.ToClientTimeByLocation(leafTime.Now(), leafTime.WITA)
witTime, err := leafTime.ToClientTimeByLocation(leafTime.Now(), leafTime.WIT)
```

To convert time using location string ([IANA](https://www.iana.org/time-zones))
```go
amsterdamTime, err := leafTime.ToClientTimeByLocationString(leafTime.Now(), "Europe/Amsterdam")
tokyoTime, err := leafTime.ToClientTimeByLocationString(leafTime.Now(), "Asia/Tokyo")
```
