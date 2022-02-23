package leafLogger

import "fmt"

type (
	MaskedEncoder map[string]Masked
	// if Aliasing is filled will replace all value with Aliasing value
	// Fill StartIdx and EndIdx with 0 to masked all value with Pattern
	// example: { Key: Value }
	// Aliasing: ***, Pattern: *, StartIdx: 0, EndIdx: 0 -> { Key: *** }
	// Pattern: *, Skipper.First: 0, Skipper.Last: 0 -> { Key: ***** }
	// Pattern: -, Skipper.First: 1, Skipper.Last: 2 -> { Key: V--ue }
	// Pattern: *, Skipper.First: 3, Skipper.Last: 3 -> { Key: Value }
	Masked struct {
		Key      string
		Aliasing string
		Pattern  string
		Skipper  Skipper
	}
	Skipper struct {
		First int
		Last  int
	}
)

func (masked Masked) encode(key string, value interface{}) string {
	var first = masked.Skipper.First
	data := fmt.Sprintf("%v", value)
	if masked.Key != key {
		return data
	}

	if masked.Aliasing != "" {
		return masked.Aliasing
	}

	lenData := len(data)
	if first > lenData {
		first = lenData
	}
	var last = lenData - masked.Skipper.Last
	if last < first {
		return data
	}

	encoded := data[:first]
	numberPattern := last - first
	if numberPattern <= 0 {
		numberPattern = lenData
	}
	for i := 0; i < numberPattern; i++ {
		encoded += masked.Pattern
	}

	encoded += data[last:]
	return encoded
}

func (masked MaskedEncoder) Encode(key string, data interface{}) interface{} {
	if mapData, ok := data.(map[string]interface{}); !ok {
		if _, exist := masked[key]; exist {
			data = masked[key].encode(key, data)
		}
	} else {
		for key, data := range mapData {
			mapData[key] = masked.Encode(key, data)
		}
		data = mapData
	}
	return data
}

