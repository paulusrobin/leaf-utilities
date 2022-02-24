package leafTypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type NullInt32 struct {
	int32 int32
	valid bool
}

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.int32)
	}
	return json.Marshal(nil)
}

func (n *NullInt32) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.int32)
	if err == nil {
		n.valid = true
	}
	return err
}

func (n NullInt32) MarshalBinary() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.int32)
	}
	return json.Marshal(nil)
}

func (n *NullInt32) UnmarshalBinary(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.int32)
	if err == nil {
		n.valid = true
	}
	return err
}

// Scan implements the Scanner interface.
func (n *NullInt32) Scan(value interface{}) error {
	sqlData := &sql.NullInt32{}
	err := sqlData.Scan(value)
	n.int32 = sqlData.Int32
	n.valid = sqlData.Valid
	return err
}

// Value implements the driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.int32, nil
}

func (n NullInt32) Val() *int32 {
	if n.valid {
		return &n.int32
	}
	return nil
}

func (n NullInt32) Valid() bool {
	return n.valid
}

func NewNullInt32(val int32) NullInt32 {
	return NullInt32{
		int32: val,
		valid: true,
	}
}
