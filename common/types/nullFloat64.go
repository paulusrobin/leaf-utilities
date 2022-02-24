package leafTypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type NullFloat64 struct {
	float64 float64
	valid   bool
}

func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.float64)
	}
	return json.Marshal(nil)
}

func (n *NullFloat64) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.float64)
	if err == nil {
		n.valid = true
	}
	return err
}

func (n NullFloat64) MarshalBinary() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.float64)
	}
	return json.Marshal(nil)
}

func (n *NullFloat64) UnmarshalBinary(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.float64)
	if err == nil {
		n.valid = true
	}
	return err
}

// Scan implements the Scanner interface.
func (n *NullFloat64) Scan(value interface{}) error {
	sqlData := &sql.NullFloat64{}
	err := sqlData.Scan(value)
	n.float64 = sqlData.Float64
	n.valid = sqlData.Valid
	return err
}

// Value implements the driver Valuer interface.
func (n NullFloat64) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.float64, nil
}

func (n NullFloat64) Val() *float64 {
	if n.valid {
		return &n.float64
	}
	return nil
}

func (n NullFloat64) Valid() bool {
	return n.valid
}

func NewNullFloat64(val float64) NullFloat64 {
	return NullFloat64{
		float64: val,
		valid:   true,
	}
}
