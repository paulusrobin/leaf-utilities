package leafTypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type NullBool struct {
	bool  bool
	valid bool
}

func (n NullBool) MarshalJSON() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.bool)
	}
	return json.Marshal(nil)
}

func (n *NullBool) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.bool)
	if err == nil {
		n.valid = true
	}
	return err
}

func (n NullBool) MarshalBinary() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.bool)
	}
	return json.Marshal(nil)
}

func (n *NullBool) UnmarshalBinary(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.bool)
	if err == nil {
		n.valid = true
	}
	return err
}

// Scan implements the Scanner interface.
func (n *NullBool) Scan(value interface{}) error {
	sqlData := &sql.NullBool{}
	err := sqlData.Scan(value)
	n.bool = sqlData.Bool
	n.valid = sqlData.Valid
	return err
}

// Value implements the driver Valuer interface.
func (n NullBool) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.bool, nil
}

func (n NullBool) Val() *bool {
	if n.valid {
		return &n.bool
	}
	return nil
}

func (n NullBool) Valid() bool {
	return n.valid
}

func NewNullBool(val bool) NullBool {
	return NullBool{
		bool:  val,
		valid: true,
	}
}
