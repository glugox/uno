package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// MarshalOnlyModel uses model and slice of field names and returns
// a json that have only those requested fields
func MarshalOnlyModel(m Model, fCol *FieldCol) ([]byte, error) {

	// Filtered copy
	c := make(map[string]interface{}, fCol.Size())

	// Orig
	reflVal := reflect.ValueOf(m).Elem()

	for _, f := range fCol.items {
		result := reflVal.FieldByName(f.Name).Interface()
		c[f.DbFieldName] = result
	}
	return json.Marshal(c)
}

func MarshalOnlyModelDefeult(m Model) ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	return MarshalOnlyModel(m, Fields(m))
}

// MarshalOnlyCollection is same as MarshalOnlyModel, but for collection
func MarshalOnlyCollection(col Collection, fCol *FieldCol) ([]byte, error) {

	fmt.Printf("fNames : %v", fCol)
	a := make([]map[string]interface{}, col.Size())
	i := 0
	for _, item := range col.Items() {

		// Filtered copy
		c := make(map[string]interface{}, fCol.Size())

		// Orig
		reflVal := reflect.ValueOf(item).Elem()

		for _, f := range fCol.Items() {
			result := reflVal.FieldByName(f.Name).Interface()
			c[f.DbFieldName] = result
		}

		a[i] = c
		i++

	}

	return json.Marshal(a)
}
