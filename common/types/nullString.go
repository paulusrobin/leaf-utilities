package leafTypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type NullString struct {
	string string
	valid  bool
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.string)
	}
	return json.Marshal(nil)
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.string)
	if err == nil {
		n.valid = true
	}
	return err
}

func (n NullString) MarshalBinary() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.string)
	}
	return json.Marshal(nil)
}

func (n *NullString) UnmarshalBinary(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.string)
	if err == nil {
		n.valid = true
	}
	return err
}

// Scan implements the Scanner interface.
func (n *NullString) Scan(value interface{}) error {
	sqlData := &sql.NullString{}
	err := sqlData.Scan(value)
	n.string = sqlData.String
	n.valid = sqlData.Valid
	return err
}

// Value implements the driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.string, nil
}

func (n NullString) Val() *string {
	if n.valid {
		return &n.string
	}
	return nil
}

func (n NullString) Valid() bool {
	return n.valid
}

func NewNullString(string string) NullString {
	return NullString{
		string: string,
		valid:  true,
	}
}
