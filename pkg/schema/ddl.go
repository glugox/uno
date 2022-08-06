package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/glugox/uno/pkg/log"
)

// DDL is a tool to extract DB information
// from our registered structs as models
type DDL struct {
	Schema *Schema
}

// NewDDL creates new DDL
func NewDDL() *DDL {
	return &DDL{
		Schema: NewSchema(),
	}
}

// Schema holds all information about our
// struct extracted database structure
type Schema struct {
	Tables *TableCol
	Logger log.Logger
}

// NewSchema creates new instance of Schema
func NewSchema() *Schema {
	logger := log.DefaultLogFactory().NewLogger()
	return &Schema{
		Tables: NewTableCol(),
		Logger: logger,
	}
}

// Configure allows to add structs
func (ddl *DDL) Configure(models ...Model) error {

	// LOOP 1: Crete tables in first loop
	for _, m := range models {
		rPtr := reflect.ValueOf(m)
		rVal := reflect.Indirect(rPtr)
		t := rVal.Type()

		// Parse struct name ( CamelCase ) => table name as snake_case
		s := fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
		ddl.Schema.Logger.Debug("DDL: Adding sruct [%s] ...", s)

		// Create and store new table
		//tbl := NewTable(t.PkgPath(), m.Meta().Name, t.Name(), rVal.Type().String(), m)
		tbl := NewTable(rVal, m)
		ddl.Schema.Tables.Append(tbl)
	}

	err := ddl.Schema.parseRelations()
	if err != nil {
		return err
	}

	//ddl.Schema.Print()
	return nil

}

// Read reads metadata tags from structs and creates schema
func (ddl *DDL) Read() *Schema {
	return ddl.Schema
}

func (s *Schema) TablesCount() int {
	return s.Tables.Size()
}

// Print returns json string representation of our schema
func (s *Schema) Dump() (string, error) {
	json, err := json.MarshalIndent(s.Tables, "", " ")
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// Print prints the dumped json to the screen
func (s *Schema) Print() {
	json, err := s.Dump()
	if err != nil {
		s.Logger.Debug("could not dump the schema as JSON. Err: %s", err)
	}
	fmt.Print(json)
}

// parseRelation checks wether the fiald is actually a rlation
//  by looking the Field's type.
// If it is an slice of some other type, it must be relation
// The function return (false, nil) if it is not a relation
func (s *Schema) parseRelations() error {

	s.Logger.Debug("parseRelations of %d tables", s.Tables.Size())

	// LOOP 2: Crete relations in 2nd loop
	tables := s.Tables.Items()
	for _, t := range tables {
		//fmt.Printf("Table: %s (%s) \n", t.Name, t.StructName)
		for _, f := range t.Fields.Items() {
			//fmt.Printf(" - %s (%s) \n", f.Name, f.Type)

			if strings.HasPrefix(f.Type, "[]") {
				rel := NewRelation(f.Name)
				// Referent field type

				refFType := strings.TrimLeft(strings.Split(f.Type, ".")[1], "[]")
				fefTable, err := s.Tables.ByStructName(refFType)
				if err != nil {
					return err
				}
				rel.Type = OneToMany
				rel.Table = fefTable.Name
				rel.Field = "id"

				t.AddRelation(rel)
			}

		}

	}

	return nil
}
