package binary

import "github.com/kelindar/binary"

func Marshal(v interface{}) (data []byte, err error) {
	return binary.Marshal(v)
}
