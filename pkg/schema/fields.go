package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/glugox/uno/pkg/utils"
	"golang.org/x/exp/slices"
)

type ObjectId int64

// Field represents metadata about
// one fild in our model struct
type Field struct {
	Name        string
	DbFieldName string
	Rel         string
	Type        string // reflect.Type
}

// FieldCol is collection of Fields
type FieldCol struct {
	items []*Field
}

// NewFieldCol returns new instance of FieldCol
func NewFieldCol() *FieldCol {
	return &FieldCol{}
}

// Append adds one Field item to the fields collection
func (col *FieldCol) Append(f *Field) {
	col.items = append(col.items, f)
}

// Items
func (col *FieldCol) Items() []*Field {
	return col.items
}

// Filter takes slice of field names and returns the same FieldCol,
// but with new Field items (filtered)
func (col *FieldCol) Filter(fNames []string) *FieldCol {
	newItems := []*Field{}
	for _, f := range col.items {
		if slices.Contains(fNames, f.Name) {
			newItems = append(newItems, f)
		}
	}
	col.items = newItems
	return col
}

// ByName looks up for a field by passed name and reduurns it
// and an error if it was not found
func (col *FieldCol) ByName(name string) (*Field, error) {
	for _, f := range col.items {
		if f.Name == name {
			return f, nil
		}
	}
	return nil, fmt.Errorf("could not find Field by name: %s", name)
}

// StringNames returns slice of strings that
// represnts field names in thie struct
func (col *FieldCol) StringNames() []string {
	names := []string{}
	for _, f := range col.items {
		//if f.Rel == "" {
		names = append(names, f.Name)
		//}
	}
	return names
}

// DBStringNames returns slice of strings that
// represnts field names in DB
func (col *FieldCol) DBStringNames() []string {
	names := []string{}
	for _, f := range col.items {
		// Dont put any relation fields here like "items", "children", etc
		if f.Rel == "" {
			names = append(names, f.DbFieldName)
		}
	}
	return names
}

// ToSqlString returns string of DB field names joined by comma (, )
func (col *FieldCol) ToSqlString() string {
	return strings.Join(col.DBStringNames(), ", ")
}

// MapDBToStruct returns a map of database feild names to struct field names
func (col *FieldCol) MapDBToStruct() map[string]string {
	m := make(map[string]string)
	for _, f := range col.items {
		m[f.DbFieldName] = f.Name
	}
	return m
}

// MapStructToDB returns a map of struct feild names to DB field names
func (col *FieldCol) MapStructToDB() map[string]string {
	m := make(map[string]string)
	for _, f := range col.items {
		m[f.Name] = f.DbFieldName
	}
	return m
}

// Size
func (col *FieldCol) Size() int {
	return len(col.items)
}

// At returns field item At specified index
func (col *FieldCol) At(index int) (*Field, error) {
	if col.Size() > index {
		return col.items[index], nil
	}
	return nil, fmt.Errorf("could not find Field at index: %d", index)
}

// MarshalJSON
func (col *FieldCol) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["items"] = col.items
	return json.Marshal(m)
}

// Fields uses reflection to parse our model struct fields and tags
// and builds a nice FieldCol that is easier to use
func Fields(m Model) *FieldCol {

	if m == nil {
		panic("can not find fields of nil model")
	}
	// Reflection Value of our  model
	refl := reflect.ValueOf(m).Elem()
	// Number of fields
	fNum := refl.NumField()
	fCol := NewFieldCol()

	// Per item vals
	var fName string
	var fRefl reflect.StructField

	for i := 0; i < fNum; i++ {
		fRefl = refl.Type().Field(i)
		fName = fRefl.Name

		//fmt.Printf(" - %d. %s \n", i, fName)

		// Try to get db as tag, and if not there...
		dbFN := fRefl.Tag.Get("db")
		if dbFN == "" {
			// ... use json tag name (split e.g. "label,omitempty")
			dbFN = strings.Split(fRefl.Tag.Get("json"), ",")[0]
		}
		if dbFN == "" || dbFN == "-" {
			// Set database field name as original struct name for the start
			dbFN = utils.ToSnakeCase(fName)
		}

		// Check for relations:
		relFN := fRefl.Tag.Get("rel")

		field := &Field{
			Name:        fName,
			DbFieldName: dbFN,
			Rel:         relFN,
			Type:        fRefl.Type.String(),
		}
		fCol.Append(field)

	}
	return fCol
}

// Fields uses reflection to parse our model struct fields and tags
// and builds a nice FieldCol that is easier to use
func QueryFields(m Model, qFields []string) *FieldCol {
	fCol := Fields(m)
	return fCol.Filter(qFields)
}

// BaseFieldNames returns slice of strings that represents all
// fields that should be loaded from database, when we are not specific
// about what to load (default)
func BaseFieldNames(m Model) []string {
	// By default get all (*) TODO?
	return Fields(m).StringNames()
}
