package influxmetric

import (
	"fmt"
	"strings"
)

// Metric represents an influx metric as used in line protocol.
type Metric struct {
	Measurement string
	TagSet      map[string]string
	FieldSet    map[string]interface{}
	UnixTime    int64
}

// Metrics represents a slice of Metric.
type Metrics []Metric

var (
	// rTagKeysTagValuesFieldKeys escapes tag keys, tag values and field keys.
	rTagKeysTagValuesFieldKeys = strings.NewReplacer(",", "\\,",
		"=", "\\=",
		" ", "\\ ")
	// rMeasurements escapes measurements.
	rMeasurements = strings.NewReplacer(",", "\\,",
		" ", "\\ ")
	// rStringFieldValues escapes string field values.
	rStringFieldValues = strings.NewReplacer("\"", "\\\"")
)

// String satisfies stringer interface.
func (m Metric) String() string {
	// add measurement
	s := fmt.Sprintf("%s", rMeasurements.Replace(m.Measurement))

	// add tag set, if exists
	if len(m.TagSet) > 0 {
		for k, v := range m.TagSet {
			// handle special keywords
			if k == "time" {
				// skip it, tags are optional
				continue
			}

			s += fmt.Sprintf(",%s=%s", rTagKeysTagValuesFieldKeys.Replace(k), rTagKeysTagValuesFieldKeys.Replace(v))
		}
	}

	// add field set
	if len(m.FieldSet) == 0 {
		// empty field set not allowed
		return ""
	}
	// set prefix for first field to space
	prefix := " "
	for k, v := range m.FieldSet {
		k = rTagKeysTagValuesFieldKeys.Replace(k)

		// handle special keywords
		if k == "time" {
			// if there are other fields, continue
			// there has to be at least one
			if len(m.FieldSet) > 1 {
				continue
			} else {
				return ""
			}
		}

		// format depends on type of value
		switch v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			s += fmt.Sprintf("%s%s=%di", prefix, k, v)
		case float32, float64:
			s += fmt.Sprintf("%s%s=%.2f", prefix, k, v)
		case bool:
			s += fmt.Sprintf("%s%s=%t", prefix, k, v)
		case string:
			s += fmt.Sprintf(`%s%s="%s"`, prefix, k,
				rStringFieldValues.Replace(v.(string)))
		default:
			// could not handle format
			return ""
		}

		// set prefix for following fields to comma
		prefix = ","
	}

	// add timestamp
	if m.UnixTime > 0 {
		s += fmt.Sprintf(" %d", m.UnixTime)
	}

	return s
}
