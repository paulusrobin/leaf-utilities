package binary

import (
	"bytes"
	"encoding/binary"
)

func UnmarshalBinary(data []byte, v interface{}) error {
	return binary.Read(bytes.NewBuffer(data), binary.BigEndian, &v)
}
