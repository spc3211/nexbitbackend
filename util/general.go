package util

import (
	"math"
	"reflect"
)

// Function to round all float64 fields in a struct to 3 decimal places
func RoundFloatFields(i interface{}) {
	v := reflect.ValueOf(i).Elem()

	for j := 0; j < v.NumField(); j++ {
		field := v.Field(j)

		if field.Kind() == reflect.Float64 {
			rounded := math.Round(field.Float()*1000) / 1000
			field.SetFloat(rounded)
		}
	}
}
