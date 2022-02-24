package leafTypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	"github.com/paulusrobin/leaf-utilities/time"
	"time"
)

type NullTime struct {
	time  time.Time
	valid bool
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.time)
	}
	return json.Marshal(nil)
}

func (n *NullTime) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.time)
	if err == nil {
		n.valid = true
	}
	return err
}

func (n NullTime) MarshalBinary() ([]byte, error) {
	if n.valid {
		return json.Marshal(n.time)
	}
	return json.Marshal(nil)
}

func (n *NullTime) UnmarshalBinary(b []byte) error {
	// Ignore null, like in the default time/time.go package.
	if string(b) == "null" {
		return nil
	}

	err := json.Unmarshal(b, &n.time)
	if err == nil {
		n.valid = true
	}
	return err
}

// Scan implements the Scanner interface.
func (n *NullTime) Scan(value interface{}) error {
	sqlData := &sql.NullTime{}
	err := sqlData.Scan(value)
	n.time = sqlData.Time
	n.valid = sqlData.Valid
	return err
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.time, nil
}

func (n NullTime) Val() *time.Time {
	if n.valid {
		return &n.time
	}
	return nil
}

func (n NullTime) Valid() bool {
	return n.valid
}

func NewNullTime() NullTime {
	return NullTime{
		time:  leafTime.Now(),
		valid: true,
	}
}

func NewNullTimeFromTime(t time.Time) NullTime {
	return NullTime{
		time:  t,
		valid: true,
	}
}
