package db

import (
	"fmt"

	"github.com/glugox/uno/pkg/schema"
)

type Adapter interface {

	// For default behaviour, see DefaultAdapter.Open
	Open(string) error

	// For default behaviour, see DefaultAdapter.Exec
	Exec(string, ...any) error

	// For default behaviour, see DefaultAdapter.Rows
	Rows(string, ...any) error

	// ScanRow takes the first arg as object
	// and tries to scan Row values from database source into it.
	ScanModel(schema.Model, *schema.Query) (mOut schema.Model, err error)

	// ScanRows takes the first arg as objects
	// and tries to scan Rows from database source into it
	// based on the Query we passed
	// E.g.
	// var menuCol MenuCollection
	// ScanRows(menuCol, NewQuery("menu", menuCol.BaseFIelds()))
	ScanCollection(schema.Collection, *schema.Query) error

	// ScanRelations looks for any given relations set in Query,
	// and queries the DB fro those and than populates into the Collection
	ScanRelations(*schema.Schema, schema.Collection, *schema.Query) error

	// Values TODO Generics (to replace "string")!
	Values(string, ...any) ([]string, error)

	// Type is one of DBAdapterMySql, DBAdapterSqlite, etc
	Type() string

	// InitMigrations creates migrations table in DB
	InitMigrations(table string) error
}

// AdapterFactory creates new Adapter implementation based
// on adapter name (mysql, sqlite, etc)
func AdapterFactory(name string) (Adapter, error) {
	switch name {
	// Sqlite
	case schema.DBAdapterSqlite:
		return NewSqliteAdapter(), nil
		// Sqlite
	case schema.DBAdapterMySql:
		return NewMySqlAdapter(), nil
	// Mysql
	default:
		return nil, fmt.Errorf("the adapter %q does not exist", name)
	}
}
