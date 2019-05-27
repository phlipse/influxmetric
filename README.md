# influxmetric

[![GoDoc](https://godoc.org/github.com/phlipse/influxmetric?status.svg)](https://godoc.org/github.com/phlipse/influxmetric)
[![Go Report Card](https://goreportcard.com/badge/github.com/phlipse/influxmetric)](https://goreportcard.com/report/github.com/phlipse/influxmetric)

influxmetric package generates metrics of type influx for use with telegraf exec input plugin.

## Usage
For examples look at *examples_test.go* file in this repository.

## Line Protocol
InfluxDB uses so called [line protocol](https://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/) to receive and store metrics.

e.g.:

```
weather,location=us-midwest temperature=82 1465839830100400200
  |    -------------------- --------------  |
  |             |             |             |
  |             |             |             |
+-----------+--------+-+---------+-+---------+
|measurement|,tag_set| |field_set| |timestamp|
+-----------+--------+-+---------+-+---------+
```

### Rules
*All rules are implemented by this package. It respects them for you!*

#### General Rules

* There must be a comma and no space between measurement and tag set.
* There must be a space between tag set and field set.
* Tag set is optional.
* At least one field set is required.
* Timestamp is optional.

#### Data Type Rules

##### Timestamp

* The minimum valid timestamp is -9223372036854775806 or 1677-09-21T00:12:43.145224194Z.
* The maximum valid timestamp is 9223372036854775806 or 2262-04-11T23:47:16.854775806Z.
* InfluxDB assumes that timestamps have nanosecond precision. --> If time is in unix seconds then append the character s to the timestamp.

##### Field Values

* Within a measurement, a field's type cannot differ within a shard, but it can differ across shards.
* InfluxDB assumes all numerical field values are floats.
* Append an i to the field value to tell InfluxDB to store the number as an integer.
* Double quote string field values, see also section "Quoting Rules".
* Specify TRUE with t, T, true, True, or TRUE.
* Specify FALSE with f, F, false, False, or FALSE.

#### Special Characters Rules

##### Tag keys, tag values and field keys
Use backslash to escape the following characters:

* commas
* equal signs
* spaces

##### Measurements
Use backslash to escape the following characters:

* commas
* spaces

##### String field values
Use backslash to escape the following characters:

* double quotes

##### Other special characters

* Line Protocol does not require users to escape the backslash character \ but will not complain if you do.
* All other special characters also do not require escaping. For example, Line Protocol handles emojis with no problem.

#### Keyword Rules

* The keyword time cannot be a field key or tag key.

#### Quoting Rules

* Never double or single quote the timestamp.
* Never single quote field values, even if they’re strings.
* Do not double or single quote measurement names, tag keys, tag values, and field keys. It is valid Line Protocol but InfluxDB assumes that the quotes are part of the name.
* Do not double quote field values that are floats, integers, or booleans. InfluxDB will assume that those values are strings.
* Do double quote field values that are strings.

## Nice To Know

* If metrics are printed out through built-in String() method, floats get truncated to two decimal places.
* Use the function ExtractValue() to extract integers, floats or booleans from strings. Numbers could be more easy graphed than strings.
* Integer values get converted to integer values in metrics. If you want to use floats, append a decimal zero. --> e.g. 8 is converted to 8i (integer not float!), if you want a float use 8.0
* This package handles errors silently and always tries to work on. There are no errors printed to STDOUT.
* A timestamp of 0 is interpreted as not set.

## License
[Apache License 2.0](https://github.com/phlipse/influxmetric/blob/master/LICENSE)
