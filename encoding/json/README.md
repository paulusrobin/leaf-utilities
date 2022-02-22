# Encoding/Binary

## JSON Serializer for Go
Dependency:
```
github.com/json-iterator/go
```
A high-performance 100% compatible drop-in replacement of "encoding/json"

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/encoding/json"
    ...
)
```

To serialize a message, simply Marshal:
```go
v := &message{
    Name:      "Roman",
    Timestamp: 1242345235,
    Payload:   []byte("hi"),
    Ssid:      []uint32{1, 2, 3},
}
encoded, err := json.Marshal(v)
```

To deserialize, Unmarshal:
```go
var v message
err := json.Unmarshal(encoded, &v)
```