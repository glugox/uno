package db

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/glugox/uno/pkg/log"
	"github.com/glugox/uno/pkg/schema"
)

// BaseAdapter
type BaseAdapter struct {
	Name   string
	DB     *sql.DB
	Logger log.Logger
}

func NewBaseAdapter(parent string) *BaseAdapter {
	return &BaseAdapter{
		Name:   parent,
		Logger: log.DefaultLogFactory().NewLogger(),
	}
}

// Type returns the type name of the adapter: "mysql", "sqlite", etc
func (o *BaseAdapter) GetMigrationSQL() string {
	return `CREATE TABLE IF NOT EXISTS migrations (id INTEGER PRIMARY KEY, name  TEXT, batch INTEGER);`
}

// Type returns the type name of the adapter: "mysql", "sqlite", etc
func (o *BaseAdapter) Type() string {
	return o.Name
}

// Open implements Adatpter interface
func (o *BaseAdapter) Open(dsn string) error {
	o.Logger.Debug("%s: connecting to the database...", o.Name)
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
func (o *BaseAdapter) Exec(q string, bind ...any) (err error) {
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
func (o *BaseAdapter) Rows(q string, bind ...any) (err error) {
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
func (o *BaseAdapter) ScanModel(m schema.Model, query *schema.Query) (mOut schema.Model, err error) {
	strSql := query.ToSQL()
	o.Logger.Verbose("SQL: %s", strSql)

	// Create collection , but later we will have only one item and take it back
	col := schema.NewCollection(m)

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
func (o *BaseAdapter) ScanCollection(col schema.Collection, query *schema.Query) (err error) {

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
func (o *BaseAdapter) ScanRelations(sch *schema.Schema, col schema.Collection, query *schema.Query) (err error) {

	fmt.Printf("ScanRelations q: %s \n", query.ToSQL())

	// Find colection's model type
	dummyRfl := col.ModelReflect()
	dummyRfl = reflect.Indirect(dummyRfl)
	t := dummyRfl.Type()

	// Get collection's model table
	tbl, err := sch.Tables.ByStructType(t.String())
	if err != nil {
		return err
	}

	// Loop all relations of the collection model's table
	for _, r := range tbl.Relations.Items() {
		//relQuery := NewQuery(f.Rel, []string{"Id"})
		relTbl, err := sch.Tables.ByName(r.Table)
		if err != nil {
			return err
		}

		relModel := reflect.New(relTbl.Reflection.Type()).Interface().(schema.Model)
		relCol := schema.NewCollection(relModel)

		var itemId schema.ObjectId

		// For al of the collection items, we need to query
		// appropriate relation items for each row:
		for _, colItem := range col.Items() {

			switch r.Type {
			case schema.OneToMany:

				// Get the reflect of the collection item that we need to set loaded relations to
				colItemRef := reflect.ValueOf(colItem)

				colItemRefElem := colItemRef.Elem()
				colItemF := colItemRefElem.FieldByName(r.Name)

				// Get primary field value (id)
				colItemPrimaryIdF := colItemRefElem.FieldByName("Id")
				itemId = schema.ObjectId(colItemPrimaryIdF.Int())

				// Build query for item relation
				relItemsQuery := schema.NewQuery(relTbl.Name, relTbl.Fields)
				relItemsQuery.AddWhere(tbl.ForeignKeyFN, fmt.Sprintf("%d", itemId))

				fmt.Printf("Rel: %+v", r)

				err = o.ScanCollection(relCol, relItemsQuery)
				if err != nil {
					return err
				}

				if colItemF.IsValid() {
					if colItemF.CanSet() {
						expectedTypeSlice := reflect.MakeSlice(colItemF.Type(), relCol.Size(), relCol.Size())

						// Set each item from queried colllection result to the new typed slice
						for relColIdx, relColItem := range relCol.Items() {
							expectedTypeSlice.Index(relColIdx).Set(reflect.ValueOf(relColItem).Elem())
						}
						//colItemF.Set(reflect.ValueOf(typedRelCol))
						colItemF.Set(expectedTypeSlice)

					} else {
						panic("Can not set rel field value!")
					}

				} else {
					panic("rel field value not valid!")
				}

			}
		}

	}

	return nil
}

// Values implements Adatpter interface. It returns a slice of values
// defined in SQL. (SELECT name FROM users) will return slice of all user names
// TODO: Generics!
func (o *BaseAdapter) Values(q string, bind ...any) (ss []string, err error) {
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
func (o *BaseAdapter) InitMigrations(table string) error {
	err := o.Exec(o.GetMigrationSQL())
	return err
}

// Bind sql arguments int string
func (o *BaseAdapter) bind(sql string, bind ...any) string {
	sql = fmt.Sprintf(sql, bind...)
	return sql
}
