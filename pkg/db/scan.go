package db

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/glugox/uno/pkg/schema"
)

// Scan function takes result set from the database (rows)
// and maps all fields to the passed Model Collection
func Scan(rows *sql.Rows, col schema.Collection, query *schema.Query) error {
	// Columns that are requested by SQL's SELECT statement
	// Database field names e.g. ["id", "parent_id", "name"]
	cols, _ := rows.Columns()

	// FieldCol object for easily manipulationg metadata
	fields := query.Fields

	// Used to create new instances of the model
	creatorRfl := col.ModelReflect()

	// Create map of field => value in each iteration so we know what to assign to each field
	var valMap map[string]interface{}
	for rows.Next() {
		columns := make([]interface{}, len(cols))

		// sql.Rows only allows scanning the data into pointers,
		// so create separate slice of pointers
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan database results into pointers slice
		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		// map database results from pointers slice to value map
		valMap = make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			valMap[colName] = *val
		}

		// Create new zero val Model, so we can scan into it
		// and append it to collection
		rValP := col.CreateModel()
		// Get the value instead pointer
		rVal := reflect.Indirect(rValP)

		t := creatorRfl.Type()
		sfNameIndex := 0

		for i := 0; i < rVal.NumField(); i++ {

			rValF := rVal.Field(i)
			sfName := rVal.Type().Field(i).Name

			field, err := fields.ByName(sfName)
			if err != nil {
				// Here we won't return an error because it is usuall that all struct
				// fields are not found in our filtered collection "fields", it is safe to continue program
				continue
			}

			// Database field name e.g. "parent_id"
			dbFName := field.DbFieldName

			// If we have the value, set it
			if item, ok := valMap[dbFName]; ok {
				setReflValue(t, rValF, item)
			}

			sfNameIndex++
		}

		col.AppendValue(rVal)

	}

	return nil
}

func FillRelation(rValF reflect.Value, relName string) {
	fmt.Printf("Fill relation %s \n", relName)
}

// setReflValue sets the value for particular field in the model struct
func setReflValue(rType reflect.Type, rValF reflect.Value, item interface{}) {
	if rValF.CanSet() {
		if item != nil {
			switch rValF.Kind() {
			// String
			case reflect.String:
				rValF.SetString(item.(string))
			// Float64
			case reflect.Float32, reflect.Float64:
				rValF.SetFloat(item.(float64))
			// Ptr
			case reflect.Ptr:
				if reflect.ValueOf(item).Kind() == reflect.Bool {
					itemBool := item.(bool)
					rValF.Set(reflect.ValueOf(&itemBool))
				}
			// Struct
			case reflect.Struct:
				rValF.Set(reflect.ValueOf(item))
			// default
			default:
				fmt.Println(rValF.Kind(), " ? >> ", reflect.ValueOf(item).Kind())
			}
		}
	}
}
