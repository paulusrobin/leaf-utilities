package leafTime

import "time"

var (
	isFreeze        bool = false
	manipulatedTime *time.Time
)

func Now() time.Time {
	if isFreeze {
		return *manipulatedTime
	}
	return time.Now()
}

func Mock(data time.Time) {
	isFreeze = true
	manipulatedTime = &data
}

func ResetMock() {
	isFreeze = false
	manipulatedTime = nil
}
