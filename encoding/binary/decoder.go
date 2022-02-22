package binary

import (
	"github.com/kelindar/binary"
)

func Unmarshal(data []byte, v interface{}) error {
	return binary.Unmarshal(data, v)
}
