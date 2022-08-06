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
	// Models is the holder for all models that we want to configue
	// along with default code models that are configured automativally
	Models  []schema.Model
	Seeders []DBSeeder
}

type DBSeeder func(*DB) error

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

// RegisterModels allows custom models registering that will be
// part of complete db Schema along with core models.
func (o *DB) RegisterModels(models []schema.Model) {
	o.Models = append(o.Models, models...)
}

// RegisterSeeder keeps track aff all handler functions that we
// need to call with DB object, so they can all seed thir parts to the DB
func (o *DB) RegisterSeeder(f DBSeeder) {
	o.Seeders = append(o.Seeders, f)
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
	mds := []schema.Model{&schema.User{}, &schema.Role{}, &schema.Menu{}, &schema.MenuItem{}}
	mds = append(mds, o.Models...)
	ddl.Configure(mds...)
	o.Schema = ddl.Read()

	err = o.Migrator.Init()
	if err != nil {
		return err
	}

	for _, seedH := range o.Seeders {
		err = seedH(o)
		if err != nil {
			return err
		}
	}

	return nil
}
