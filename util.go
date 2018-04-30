// +build windows

package wind

import (
	"fmt"
	"time"
)

// SectorCons return a string slice of sector constituent
func SectorCons(id string, date time.Time) ([]string, error) {
	var o = Option{
		Date:     date.Format("2006-01-02"),
		SectorID: id,
		Field:    "wind_code",
	}

	_, _, _, d, err := Wset("sectorconstituent", &o)
	if err != nil {
		return nil, err
	}

	data, ok := d.([]interface{})
	if !ok {
		return nil, fmt.Errorf("type assertion fail")
	}

	var codes = make([]string, len(data))
	for i, x := range data {
		codes[i] = x.(string)
	}

	return codes, nil
}

// Fields return the wind fields of account codes
func Fields(accCodes []string) ([]string, error) {
	fields := make([]string, len(accCodes))

	for i, accCode := range accCodes {
		var ok bool
		if fields[i], ok = AccCodeToWindField[accCode]; !ok {
			return nil, fmt.Errorf("unexpected account code: %q", accCode)
		}
	}

	return fields, nil
}
