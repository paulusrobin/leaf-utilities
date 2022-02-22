package leafMandatory

import (
	"database/sql"
	"database/sql/driver"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	"fmt"
	"strings"
)

type (
	DeviceType int
	DeviceInfo struct {
		id   DeviceType
		code string
		name string
	}
)

func (d DeviceInfo) ID() DeviceType {
	return d.id
}

func (d DeviceInfo) Name() string {
	return d.name
}

func (d DeviceInfo) Code() string {
	return d.code
}

func (d DeviceInfo) JSON() map[string]interface{} {
	return map[string]interface{}{
		"id":   d.ID(),
		"code": d.Code(),
		"name": d.Name(),
	}
}

const (
	Android DeviceType = iota + 1
	Ios
	Web
	MobileWeb
)

var deviceInfos = []DeviceInfo{
	{Android, "ANDROID", "Android"},
	{Ios, "IOS", "Ios"},
	{Web, "WEB", "Website"},
	{MobileWeb, "MWEB", "Mobile Website"},
}

func (s DeviceType) Info() DeviceInfo {
	return deviceInfos[s-1]
}

func (s DeviceType) JSON() map[string]interface{} {
	if s < 1 || int(s) > len(deviceInfos) {
		return map[string]interface{}{}
	}
	return deviceInfos[s-1].JSON()
}

func (s DeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Info().Name())
}

func (s *DeviceType) UnmarshalJSON(b []byte) error {
	var strAction = ""
	err := json.Unmarshal(b, &strAction)
	if err != nil {
		return err
	}
	*s, err = DeviceFromString(strAction)
	return err
}

// Scan implements the Scanner interface.
func (s *DeviceType) Scan(value interface{}) error {
	sqlData := &sql.NullString{}
	err := sqlData.Scan(value)
	if err != nil || !sqlData.Valid {
		return err
	}

	*s, err = DeviceFromString(sqlData.String)
	return err
}

// Value implements the driver Valuer interface.
func (s DeviceType) Value() (driver.Value, error) {
	return s.Info().Name(), nil
}

func DeviceFromString(str string) (DeviceType, error) {
	lowerStr := strings.ToLower(str)
	for i, j := 0, len(deviceInfos)-1; i <= j; i, j = i+1, j-1 {
		if strings.ToLower(deviceInfos[i].name) == lowerStr {
			return DeviceType(i + 1), nil
		}
		if strings.ToLower(deviceInfos[j].name) == lowerStr {
			return DeviceType(j + 1), nil
		}
	}
	return -1, fmt.Errorf("invalid device")
}

func DeviceFromStringCode(str string) (DeviceType, error) {
	lowerStr := strings.ToLower(str)
	for i, j := 0, len(deviceInfos)-1; i <= j; i, j = i+1, j-1 {
		if strings.ToLower(deviceInfos[i].code) == lowerStr {
			return DeviceType(i + 1), nil
		}
		if strings.ToLower(deviceInfos[j].code) == lowerStr {
			return DeviceType(j + 1), nil
		}
	}
	return -1, fmt.Errorf("invalid device")
}
