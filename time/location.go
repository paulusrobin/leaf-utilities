package leafTime

import "time"

type (
	Location struct {
		loc   time.Location
		valid bool
	}
)

func (l Location) Location() *time.Location {
	if !l.valid {
		return nil
	}
	return &l.loc
}

func (l Location) Valid() bool {
	return l.valid
}

func createLocation(loc string) Location {
	locationData, err := time.LoadLocation(loc)
	if err != nil {
		return Location{
			loc:   time.Location{},
			valid: false,
		}
	}
	return Location{
		loc:   *locationData,
		valid: true,
	}
}

/*
	====================================
	 Initialization Common Location
	====================================
*/

var (
	WIB  Location
	WITA Location
	WIT  Location
)

func init() {
	WIB = createLocation("Asia/Jakarta")
	WITA = createLocation("Asia/Makassar")
	WIT = createLocation("Asia/Jayapura")
}
