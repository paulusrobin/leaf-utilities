package leafTime

import (
	"time"
)

func ToUTCTime(clientTime time.Time) time.Time {
	return clientTime.In(time.UTC)
}

func ToServerTime(clientTime time.Time) time.Time {
	return clientTime.In(time.Local)
}

func ToClientTimeByLocation(t time.Time, clientLoc Location) (time.Time, error) {
	if !clientLoc.Valid() {
		return t, InvalidLocation
	}
	return t.In(clientLoc.Location()), nil
}

func ToClientTimeByLocationString(t time.Time, clientLoc string) (time.Time, error) {
	loc, err := time.LoadLocation(clientLoc)
	if err != nil {
		return t, err
	}
	return t.In(loc), nil
}
