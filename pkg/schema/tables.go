package schema

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/glugox/uno/pkg/utils"
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
	// Foreign Key Field Name is used in relations.
	// When we want to reference this table from another. The another
	// table should have this field name. E.g. this table is menus, than
	// in the menu_items table we shoud have menu_id field
	// that references to this table
	ForeignKeyFN string
	Relations    *RelationCol
	Reflection   reflect.Value
}

// NewTable creates new instance of Table
func NewTable(reflectVal reflect.Value, m Model) *Table {

	t := reflectVal.Type()
	tblName := TableName(m)

	return &Table{
		// e.g. "github.com/glugox/uno/pkg"
		Path: t.PkgPath(),
		// e.g. "user"
		Name: tblName,
		// e.g. "User"
		StructName: t.Name(),
		// e.g. "uno.User"
		StructType:   reflectVal.Type().String(),
		Fields:       Fields(m),
		ForeignKeyFN: fmt.Sprintf("%s_id", tblName),
		Relations:    NewRelationCol(),

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

// TableName tages a struct model (e.g. UserRole) and
// returns the snake case based on that struct name (user_role)
func TableName(m Model) string {

	if n, ok := m.(Namer); ok {
		return n.Name()
	}

	rVal := reflect.ValueOf(m)
	return utils.ToSnakeCase(rVal.Type().Elem().Name())
}

// ToString retruns string representation of the passed model
func ToString(m Model) string {
	return "[" + TableName(m) + "]"
}
