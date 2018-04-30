package wind

import (
	"reflect"
	"strconv"
)

// Option is the options provide to wind
type Option struct {
	ReportType      int    `wind:"rptType"` // 1,2,3,4
	Period          byte   `wind:"Period"`
	ReportDate      string `wind:"rptDate"`   // yyyymmdd
	TradeDate       string `wind:"tradeDate"` // yyyymmdd
	PriceAdj        string `wind:"priceAdj"`
	Cycle           string `wind:"cycle"`
	Date            string `wind:"date"` // yyyy-mm-dd
	SectorID        string `wind:"sectorid"`
	Field           string `wind:"field"`
	Days            string `wind:"Days"`
	Fill            string `wind:"Fill"`
	TradingCalendar string `wind:"TradingCalendar"`
	StartDate       string `wind:"startDate"`
	EndDate         string `wind:"endDate"`
	Unit            int    `wind:"unit"`
	Order           int    `wind:"order"`
}

func (o *Option) String() string {
	if o == nil {
		return ""
	}

	var s, sep string

	v := reflect.ValueOf(o).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if !isEffective(v.Field(i)) {
			continue
		}

		key := t.Field(i).Tag.Get("wind")
		var val string
		switch v.Field(i).Kind() {
		case reflect.Bool:
			val = strconv.FormatBool(v.Field(i).Bool())
		case reflect.Int:
			val = strconv.FormatInt(v.Field(i).Int(), 10)
		case reflect.String:
			val = v.Field(i).String()
		case reflect.Uint8: // byte
			val = string(v.Field(i).Uint())
		}
		s += sep + key + "=" + val
		sep = ";"
	}

	return s
}

// Clear all options
func (o *Option) Clear() {
	*o = Option{}
}

// if non-zero value, return true
func isEffective(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int:
		return v.Int() != 0
	case reflect.Uint8:
		return v.Uint() != 0
	case reflect.String:
		return v.String() != ""
	default:
		return false
	}
}
