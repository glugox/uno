package db

import (
	"database/sql"

	"github.com/glugox/uno/pkg/log"
	"github.com/glugox/uno/pkg/schema"
	_ "github.com/go-sql-driver/mysql"
)

// MySqlAdapter
type MySqlAdapter struct {
	Name   string
	Base   *BaseAdapter
	DB     *sql.DB
	Logger log.Logger
}

// NewMySqlAdapter
func NewMySqlAdapter() (a *MySqlAdapter) {
	logger := log.DefaultLogFactory().NewLogger()
	a = &MySqlAdapter{
		Name:   schema.DBAdapterMySql,
		Logger: logger,
		Base:   NewBaseAdapter(schema.DBAdapterMySql),
	}

	a.Logger.Debug("created new MySqlAdapter")
	return
}

// Type returns the type name of the adapter: "mysql", "sqlite", etc
func (o *MySqlAdapter) GetMigrationSQL() string {
	return `CREATE TABLE IF NOT EXISTS migrations (id INT NOT NULL AUTO_INCREMENT, name  VARCHAR(255), batch INT, PRIMARY KEY(id));`
}

// Type returns the type name of the adapter: "mysql", "sqlite", etc
func (o *MySqlAdapter) Type() string {
	return o.Name
}

// Open implements Adatpter interface
func (o *MySqlAdapter) Open(dsn string) error {
	err := o.Base.Open(dsn)
	if err != nil {
		return err
	}
	o.DB = o.Base.DB
	return nil
}

// Exec implements Adatpter interface.
// It only executes sql query
func (o *MySqlAdapter) Exec(q string, bind ...any) (err error) {
	return o.Base.Exec(q, bind...)
}

// Rows implements Adatpter interface.
// It returns rows typical for SELECT statement
func (o *MySqlAdapter) Rows(q string, bind ...any) (err error) {
	return o.Base.Rows(q, bind...)
}

// ScanModel implements db.Adapter.ScanModel
// While ScanCollection scans into the passed collection, this fnction returns
// new model instance
func (o *MySqlAdapter) ScanModel(m schema.Model, query *schema.Query) (mOut schema.Model, err error) {
	return o.Base.ScanModel(m, query)
}

// ScanRows implements db.Adapter.ScanRows
func (o *MySqlAdapter) ScanCollection(col schema.Collection, query *schema.Query) (err error) {
	return o.Base.ScanCollection(col, query)
}

// ScanRows implements db.Adapter.ScanRelations
func (o *MySqlAdapter) ScanRelations(sch *schema.Schema, col schema.Collection, query *schema.Query) (err error) {
	return o.Base.ScanRelations(sch, col, query)
}

// Values implements Adatpter interface. It returns a slice of values
// defined in SQL. (SELECT name FROM users) will return slice of all user names
// TODO: Generics!
func (o *MySqlAdapter) Values(q string, bind ...any) (ss []string, err error) {
	return o.Base.Values(q, bind...)
}

// // InitMigrations creates the migrations table in the DB
func (o *MySqlAdapter) InitMigrations(table string) error {
	// Since default one is sqlite, here we don't call base in order
	//  to get own (mysql) migration sql
	return o.Exec(o.GetMigrationSQL())
}
