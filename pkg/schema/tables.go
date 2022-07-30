package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Table is the DB representation of our struct
type Table struct {
	// With wull name prefix e.g. "github.com/glugox/uno/User"
	Path string
	// DB Name (we are in DB context)
	Name       string
	StructName string
	StructType string
	Fields     *FieldCol
	Relations  *RelationCol
	Reflection reflect.Value
}

// NewTable creates new instance of Table
// func NewTable(path, name, structName string, structType interface{}, o Model) *Table {
// 	return &Table{
// 		// e.g. "github.com/glugox/uno/pkg"
// 		Path: path,
// 		// e.g. "user"
// 		Name: name,
// 		// e.g. "User"
// 		StructName: structName,
// 		// e.g. "uno.User"
// 		StructType: structType.(string),
// 		Fields:     Fields(o),
// 		Relations:  NewRelationCol(),

// 		// Reflection
// 		Reflection: reflect.Value,
// 	}
// }

// NewTable(t.PkgPath(), m.Meta().Name, t.Name(), rVal.Type().String(), m)

func NewTable(reflectVal reflect.Value, m Model) *Table {

	t := reflectVal.Type()

	return &Table{
		// e.g. "github.com/glugox/uno/pkg"
		Path: t.PkgPath(),
		// e.g. "user"
		Name: m.Meta().Name,
		// e.g. "User"
		StructName: t.Name(),
		// e.g. "uno.User"
		StructType: reflectVal.Type().String(),
		Fields:     Fields(m),
		Relations:  NewRelationCol(),

		// Reflection
		Reflection: reflectVal,
	}
}

func (t *Table) AddRelation(r *Relation) {
	t.Relations.Append(r)
}

// TableCol is collection of Tables
type TableCol struct {
	items []*Table
}

// NewTableCol returns new instance of TableCol
func NewTableCol() *TableCol {
	return &TableCol{
		items: []*Table{},
	}
}

// Append adds one Table item to the table collection
func (col *TableCol) Append(t *Table) {
	col.items = append(col.items, t)
}

// Items
func (col *TableCol) Items() []*Table {
	return col.items
}

// Size
func (col *TableCol) Size() int {
	return len(col.items)
}

// At returns table item At specified index
func (col *TableCol) At(index int) (*Table, error) {
	if col.Size() > index {
		return col.items[index], nil
	}
	return nil, fmt.Errorf("could not find Table at index: %d", index)
}

// ByName looks up for a table by passed name and reduurns it
// and an error if it was not found
func (col *TableCol) ByName(name string) (*Table, error) {
	for _, f := range col.items {
		if f.Name == name {
			return f, nil
		}
	}
	return nil, fmt.Errorf("could not find Table by name: %s", name)
}

// ByStructName is same as ByName but looks for struct name e.g. "UserRole"
// instead table name (user_role)
func (col *TableCol) ByStructName(name string) (*Table, error) {
	for _, t := range col.items {
		if t.StructName == name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("could not find Table by struct name: %s", name)
}

// ByStructName is same as ByName but looks for struct type e.g. "uno.UserRole"
// instead table name (user_role)
func (col *TableCol) ByStructType(name string) (*Table, error) {
	for _, t := range col.items {
		if t.StructType == name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("could not find Table by struct name: %s", name)
}

// MarshalJSON
func (col *TableCol) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["items"] = col.items
	return json.Marshal(m)
}
