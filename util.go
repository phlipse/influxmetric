package influxmetric

import (
	"strconv"
)

// ExtractValue trys to extract value of type int64, float64 or boolean from given string.
func ExtractValue(s string) interface{} {
	// check conversion to int
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}

	// check conversion to float
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}

	// check conversion to bool
	b, err := strconv.ParseBool(s)
	if err == nil {
		return b
	}

	return s
}

// MustFloat converts given object to float64.
func MustFloat(i interface{}) float64 {
	switch i.(type) {
	case int:
		return float64(i.(int))
	case int8:
		return float64(i.(int8))
	case int16:
		return float64(i.(int16))
	case int32:
		return float64(i.(int32))
	case int64:
		return float64(i.(int64))
	case uint:
		return float64(i.(uint))
	case uint8:
		return float64(i.(uint8))
	case uint16:
		return float64(i.(uint16))
	case uint32:
		return float64(i.(uint32))
	case uint64:
		return float64(i.(uint64))
	case float32:
		return float64(i.(float32))
	case float64:
		return float64(i.(float64))
	default:
		return 0.0
	}
}
