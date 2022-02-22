# Encoding/Binary

## Binary Serializer for Go
Dependency:
```
github.com/kelindar/binary
```
Encode/Decode arbitrary golang data structures of variable size. This package extends support to arbitrary, variable-sized values by prefixing these values with their varint-encoded size, recursively. This was originally inspired by Alec Thomas's binary package, but I've reworked the serialization format and improved the performance and size.

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/encoding/binary"
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
encoded, err := binary.Marshal(v)
```

To deserialize, Unmarshal:
```go
var v message
err := binary.Unmarshal(encoded, &v)
```