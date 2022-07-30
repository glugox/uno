package db

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/glugox/uno/pkg/log"
	"github.com/glugox/uno/pkg/schema"
	_ "github.com/mattn/go-sqlite3"
)

// SqliteAdapter
type SqliteAdapter struct {
	Name   string
	DB     *sql.DB
	Logger log.Logger
}

// NewSqliteAdapter
func NewSqliteAdapter() (a *SqliteAdapter) {
	logger := log.DefaultLogFactory().NewLogger()
	a = &SqliteAdapter{
		Name:   schema.DBAdapterSqlite,
		Logger: logger,
	}
	a.Logger.Debug("created new SqliteAdapter")
	return
}

// Type returns the type name of the adapter: "mysql", "sqlite", etc
func (o *SqliteAdapter) Type() string {
	return o.Name
}

// Type returns the type name of the adapter: "mysql", "sqlite", etc
func (o *SqliteAdapter) GetMigrationSQL() string {
	return `CREATE TABLE IF NOT EXISTS migrations (id INTEGER PRIMARY KEY, name  TEXT, batch INTEGER);`
}

// Open implements Adatpter interface
func (o *SqliteAdapter) Open(dsn string) error {
	o.Logger.Info("%s: connecting to the database...", o.Name)
	o.Logger.Verbose("DSN: %s", dsn)
	db, err := sql.Open(o.Type(), dsn)
	if err != nil {
		return fmt.Errorf("could not open to database! Error: %s", err)
	}

	// Check if we have succesfully connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("could not connect to the database! Error: %s", err)
	}

	// All good!
	o.Logger.Success("successfully connected to the database!")

	o.DB = db
	return nil
}

// Exec implements Adatpter interface.
// It only executes sql query
func (o *SqliteAdapter) Exec(q string, bind ...any) (err error) {
	q = o.bind(q, bind...)

	o.Logger.Verbose("SQL: %s", q)
	_, err = o.DB.Exec(q)
	if err != nil {
		o.Logger.Error("Could not execute SQL. Err: %s", err)
		return
	}
	return nil
}

// Rows implements Adatpter interface.
// It returns rows typical for SELECT statement
func (o *SqliteAdapter) Rows(q string, bind ...any) (err error) {
	q = o.bind(q, bind...)

	o.Logger.Verbose("SQL: %s", q)
	rows, err := o.DB.Query(q)
	if err != nil {
		o.Logger.Error("Could not execute SQL. Err: %s", err)
		return
	}
	rows.Close()
	return nil
}

// ScanModel implements db.Adapter.ScanModel
// While ScanCollection scans into the passed collection, this fnction returns
// new model instance
func (o *SqliteAdapter) ScanModel(m schema.Model, query *schema.Query) (mOut schema.Model, err error) {
	strSql := query.ToSQL()
	o.Logger.Verbose("SQL: %s", strSql)

	// Create collection , but later we will have only one item and take it back
	col := m.Collection()

	rows, err := o.DB.Query(strSql)
	if err != nil {
		o.Logger.Error("Could not execute SQL. Err: %s", err)
		return
	}

	// Do the scanning for all of our rows
	err = Scan(rows, col, query)
	if err != nil {
		o.Logger.Error("Could not scan Rows into Model. Err: %s", err)
		return
	}

	mOut = col.Items()[0]

	rows.Close()
	return mOut, nil
}

// ScanRows implements db.Adapter.ScanRows
func (o *SqliteAdapter) ScanCollection(col schema.Collection, query *schema.Query) (err error) {

	strSql := query.ToSQL()
	o.Logger.Verbose("SQL: %s", strSql)

	rows, err := o.DB.Query(strSql)
	if err != nil {
		o.Logger.Error("Could not execute SQL. Err: %s", err)
		return
	}

	// Do the scanning for all of our rows
	err = Scan(rows, col, query)
	if err != nil {
		o.Logger.Error("Could not scan Rows into Model. Err: %s", err)
		return
	}

	rows.Close()
	return nil
}

// ScanRows implements db.Adapter.ScanRelations
func (o *SqliteAdapter) ScanRelations(sch *schema.Schema, col schema.Collection, query *schema.Query) (err error) {

	fmt.Printf("ScanRelations q: %s \n", query.ToSQL())

	dummyRfl := col.ModelReflect()
	dummyRfl = reflect.Indirect(dummyRfl)
	t := dummyRfl.Type()

	fmt.Printf("Type %s", t)

	tbl, err := sch.Tables.ByStructType(t.String())
	if err != nil {
		return err
	}
	fmt.Printf("Table %s", tbl.Name)

	for _, r := range tbl.Relations.Items() {
		fmt.Printf(" Relation: %v \n", r)
		//relQuery := NewQuery(f.Rel, []string{"Id"})
		relTbl, err := sch.Tables.ByName(r.Table)
		if err != nil {
			return err
		}

		//relModel := reflect.New(relTbl.Reflection.Type()).Interface().(Model)
		//relCol := relModel.Collection()

		switch r.Type {
		case schema.OneToMany:
			fmt.Printf("Relation table (StructType) : %s, field: %s", relTbl.StructType, r.Field)
			//relCtx := NewEntityContext(sch.DB, relModel)
			/*relModel, err := relCtx.All()
			if err != nil {
				return err
			}*/

			/*for _, relItem := range relModel {
				fmt.Printf("Rel item :%v", relItem)
			}*/
		}

	}

	// Do the scanning for all of our rows
	/*err = Scan(rows, col, query)
	if err != nil {
		o.Logger.Error("Could not scan Rows into Model. Err: %s", err)
		return
	}*/

	return nil
}

// Values implements Adatpter interface. It returns a slice of values
// defined in SQL. (SELECT name FROM users) will return slice of all user names
// TODO: Generics!
func (o *SqliteAdapter) Values(q string, bind ...any) (ss []string, err error) {
	q = o.bind(q, bind...)
	n := 0
	v := ""

	o.Logger.Verbose("SQL: %s", q)
	rows, err := o.DB.Query(q)
	if err != nil {
		return nil, err
	}

	cs, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	o.Logger.Verbose(" |_ Columns: %v", cs)

	for rows.Next() {
		rows.Scan(&v)
		o.Logger.Verbose("   |_ %d.Val: %s", n, v)
		ss = append(ss, v)
		n++
	}

	o.Logger.Verbose("Rows affected: %d", n)
	return ss, nil
}

// // InitMigrations creates the migrations table in the DB
func (o *SqliteAdapter) InitMigrations(table string) error {
	err := o.Exec(o.GetMigrationSQL())
	return err
}

// Bind sql arguments int string
func (o *SqliteAdapter) bind(sql string, bind ...any) string {
	sql = fmt.Sprintf(sql, bind...)
	return sql
}
