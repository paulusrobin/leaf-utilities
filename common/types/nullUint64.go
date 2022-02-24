package leafTypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	"fmt"
)

type NullUInt64 struct {
	int64 uint64
	valid bool
}

func (n NullUInt64) MarshalJSON() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.int64)
	}
	return json.Marshal(nil)
}

func (n *NullUInt64) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.int64)
	if err == nil && string(b) != "null" {
		n.valid = true
	}
	return err
}

func (n NullUInt64) MarshalBinary() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.int64)
	}
	return json.Marshal(nil)
}

func (n *NullUInt64) UnmarshalBinary(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.int64)
	if err == nil {
		n.valid = true
	}
	return err
}

// Scan implements the Scanner interface.
func (n *NullUInt64) Scan(value interface{}) error {
	sqlData := &sql.NullInt64{}
	err := sqlData.Scan(value)
	if sqlData.Int64 < 0 {
		return fmt.Errorf("error negative value")
	}

	n.int64 = uint64(sqlData.Int64)
	n.valid = sqlData.Valid
	return err
}

// Value implements the driver Valuer interface.
func (n NullUInt64) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.int64, nil
}

func (n NullUInt64) Val() *uint64 {
	if n.valid {
		return &n.int64
	}
	return nil
}

func (n NullUInt64) Valid() bool {
	return n.valid
}

func NewNullUInt64(val uint64) NullUInt64 {
	return NullUInt64{
		int64: val,
		valid: true,
	}
}
