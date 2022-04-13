package spanType

import (
	"fmt"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type (
	// Class .
	Class struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// Enum .
	Enum int
)

const (
	Span Enum = iota + 1
	DataStore
	MessageProducer
	External
)

var constants = []Class{
	{"Span", "Segment"},
	{"DataStore", "Data Store Segment"},
	{"MessageProducer", "Message Producer Segment"},
	{"External", "External Segment"},
}

// Code .
func (c Enum) Code() string {
	if c < 1 || int(c) > len(constants) {
		return ""
	}
	return constants[c-1].Code
}

// Name .
func (c Enum) Name() string {
	if c < 1 || int(c) > len(constants) {
		return ""
	}
	return constants[c-1].Name
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *Enum) UnmarshalParam(src string) error {
	index := findIndex(src, func(c Class) string {
		return c.Code
	})

	if index == 0 {
		return fmt.Errorf("unknown enum constant")
	}

	*c = Enum(index)
	return nil
}

// MarshalJSON presents value to the client
func (c Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *Enum) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal, func(c Class) string {
		return c.Code
	})

	if index == 0 {
		return fmt.Errorf("unknown enum constant")
	}

	*c = Enum(index)
	return nil
}

func (c *Enum) UnmarshalName(val string) error {
	index := findIndex(val, func(c Class) string {
		return c.Name
	})

	if index == 0 {
		return fmt.Errorf("unknown enum constant")
	}

	*c = Enum(index)
	return nil
}

func findIndex(code string, selector func(c Class) string) int {
	for i, v := range constants {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteConstants() []Class {
	list := make([]Class, len(constants))
	copy(list, constants)
	return list
}
