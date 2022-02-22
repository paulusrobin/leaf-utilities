package binary

import (
	"bytes"
	"encoding/binary"
)

func MarshalBinary(v interface{}) (data []byte, err error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, &v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
