package db

import (
	"errors"

	"github.com/glugox/uno/pkg/config"
	"github.com/glugox/uno/pkg/log"
	"github.com/glugox/uno/pkg/schema"
)

type DB struct {
	Config   *config.DBConfig
	Migrator Migrator
	Logger   log.Logger
	Adapter  Adapter
	Schema   *schema.Schema
}

// NewDB returns new DB instance
func NewDB() *DB {
	db := &DB{
		Adapter: NewSqliteAdapter(),
		Logger:  log.DefaultLogFactory().NewLogger(),
	}
	m := NewMigrate(db)
	db.Migrator = m

	return db
}

// WithConfig returns new DB instance with specified db config
func DBWithConfig(cfg *config.DBConfig) (*DB, error) {
	db := NewDB()
	db.Config = cfg

	// This will override the Migrator instance that
	// was set in db.NewDB
	if db.Migrator != nil {
		db.Migrator.SetConfig(cfg.Migrate)
	}

	if len(cfg.Name) > 0 && cfg.Name != db.Adapter.Type() {
		a, err := AdapterFactory(cfg.Name)
		db.Adapter = a
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// Init checks needed data and tries to connect to database
// also, tries to migrate any migrations available
func (o *DB) Init() error {
	if o.Config == nil {
		return errors.New("please specify Config for the DB")
	}

	// Try to connect to the database
	err := o.Adapter.Open(o.Config.ToString())
	if err != nil {
		return err
	}

	ddl := schema.NewDDL()
	ddl.Configure(&schema.User{}, &schema.Role{}, &schema.Menu{}, &schema.MenuItem{})
	o.Schema = ddl.Read()

	return o.Migrator.Init()
}
