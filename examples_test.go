package influxmetric_test

import (
	"fmt"
	"reflect"

	"github.com/phlipse/influxmetric"
)

func ExampleMetric() {
	var m = influxmetric.Metric{}

	m.Measurement = " W\"e,tt=e r "

	m.TagSet = make(map[string]string)
	m.TagSet["country"] = " DE "
	//m.TagSet["ci ty"] = "Würzburg"

	m.FieldSet = make(map[string]interface{})
	m.FieldSet[" temp"] = 10
	//m.FieldSet["hum"] = 48.2
	
	//m.UnixTime = time.Now().UnixNano()

	fmt.Println(m)
	// Output:
	// W_e_tt_er,country=DE temp=10i
}

func ExampleExtractValue() {
	fmt.Println(reflect.TypeOf(influxmetric.ExtractValue("-8")))
	fmt.Println(reflect.TypeOf(influxmetric.ExtractValue("42.8")))
	fmt.Println(reflect.TypeOf(influxmetric.ExtractValue("True")))
	fmt.Println(reflect.TypeOf(influxmetric.ExtractValue("true false")))

	// Output:
	// int64
	// float64
	// bool
	// string
}
