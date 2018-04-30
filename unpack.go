// +build windows

package wind

import (
	"time"
	"unsafe"

	"naturebridge-asset.com/util"
)

const (
	vtArray    = 0x2000
	vtR8       = 5
	vtDate     = 7
	vtBstr     = 8
	vtVariant  = 12
	vtTypemask = 0xfff
)

const (
	sizePtr = 8
	sizeDbl = 8
	sizeVar = 24
)

var basedate time.Time

type apiout struct {
	errorCode int32
	stateCode int32
	requestID int64
	codes     variant
	_         int64
	fields    variant
	_         int64
	times     variant
	_         int64
	data      variant
}

// pVal maybe:
//     _fields_ = [("llVal", c_int64),
//                 ("lVal", c_int32),
//                 ("bVal", c_int8),
//                 ("iVal", c_int16),
//                 ("fltVal", c_float),
//                 ("dblVal", c_double),
//                 ("date", c_double),
//                 ("bstrVal", c_wchar_p),
//                 ("pbVal", POINTER(c_byte)),
//                 ("piVal", POINTER(c_int16)),
//                 ("plVal", POINTER(c_int32)),
//                 ("pllVal", POINTER(c_int64)),
//                 ("pfltVal", POINTER(c_float)),
//                 ("pdblVal", POINTER(c_double)),
//                 ("pdate", POINTER(c_double)),
//                 ("pbstrVal", POINTER(c_wchar_p)),
//                 ("parray", POINTER(c_safearray)),
//                 ("pvarVal", POINTER(c_variant)),
//                 ("__VARIANT_NAME_4", c_tagBRECORD)]
type variant struct {
	vt   uint16
	wr1  uint16
	wr2  uint16
	wr3  uint16
	pVal uintptr
}

// pVal maybe:
//     _fields_ = [("pbVal", POINTER(c_byte)),
//                 ("piVal", POINTER(c_int16)),
//                 ("plVal", POINTER(c_int32)),
//                 ("pllVal", POINTER(c_int64)),
//                 ("pfltVal", POINTER(c_float)),
//                 ("pdblVal", POINTER(c_double)),
//                 ("pdate", POINTER(c_double)),
//                 ("pbstrVal", POINTER(c_wchar_p)),
//                 ("pvarVal", POINTER(c_variant))]
type safearray struct {
	cDims      uint16
	fFeatures  uint16
	cbElements uint32
	cLocks     uint32
	pVal       uintptr
	rgsabound  [3]safearraybound
}

type safearraybound struct {
	cElements uint32
	lLbound   int32
}

func (data *variant) getData() interface{} {
	if data.vt&vtArray == 0 {
		return data.getScala()
	}
	return data.getList()
}

func (data *variant) getScala() interface{} {
	switch data.vt & vtTypemask {
	case vtR8:
		return *(*float64)(unsafe.Pointer(&(data.pVal)))
	case vtDate:
		dur := *(*float64)(unsafe.Pointer(&(data.pVal)))
		return basedate.AddDate(0, 0, int(dur))
	case vtBstr:
		return util.UTF16PtrToString((*uint16)(unsafe.Pointer(data.pVal)))
	}
	return nil
}

func (data *variant) getList() interface{} {
	array := (*safearray)(unsafe.Pointer(data.pVal))
	count := array.getCount()
	pVal := array.pVal

	switch data.vt & vtTypemask {
	case vtR8:
		list := make([]float64, count)
		for i := 0; i < count; i++ {
			list[i] = *(*float64)(unsafe.Pointer(pVal + uintptr(sizeDbl*i)))
		}
		return list
	case vtDate:
		list := make([]time.Time, count)
		for i := 0; i < count; i++ {
			dur := *(*float64)(unsafe.Pointer(pVal + uintptr(sizeDbl*i)))
			list[i] = basedate.AddDate(0, 0, int(dur))
		}
		return list
	case vtBstr:
		list := make([]string, count)
		for i := 0; i < count; i++ {
			p := (**uint16)(unsafe.Pointer(pVal + uintptr(sizePtr*i)))
			list[i] = util.UTF16PtrToString(*p)
		}
		return list
	case vtVariant:
		list := make([]interface{}, count)
		for i := 0; i < count; i++ {
			p := (*variant)(unsafe.Pointer(pVal + uintptr(sizeVar*i)))
			list[i] = p.getScala()
		}
		return list
	}

	return nil
}

func (a *safearray) getCount() int {
	dim := a.cDims
	if dim == 0 {
		return 0
	}

	count := 1
	for i := 0; i < int(dim); i++ {
		count *= int(a.rgsabound[i].cElements)
	}
	return count
}

func init() {
	basedate = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.Local)
}
