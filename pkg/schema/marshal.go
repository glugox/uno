package schema

import (
	"encoding/json"
	"reflect"
)

// MarshalOnlyModel uses model and slice of field names and returns
// a json that have only those requested fields
func MarshalOnlyModel(m Model, fNames []string) ([]byte, error) {

	// Filtered copy
	c := make(map[string]interface{}, len(fNames))

	// Orig
	reflVal := reflect.ValueOf(m).Elem()

	for _, fName := range fNames {
		result := reflVal.FieldByName(fName).Interface()
		c[fName] = result
	}
	return json.Marshal(c)
}

// MarshalOnlyCollection is same as MarshalOnlyModel, but for collection
func MarshalOnlyCollection(col Collection, fNames []string) ([]byte, error) {

	a := make([]map[string]interface{}, col.Size())
	i := 0
	for _, item := range col.Items() {

		// Filtered copy
		c := make(map[string]interface{}, len(fNames))

		// Orig
		reflVal := reflect.ValueOf(item).Elem()

		for _, fName := range fNames {
			result := reflVal.FieldByName(fName).Interface()
			c[fName] = result
		}

		a[i] = c
		i++

	}

	return json.Marshal(a)
}
