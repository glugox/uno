package db

import (
	"os"
	"path"
	"path/filepath"

	"github.com/glugox/uno/pkg/config"
	"golang.org/x/exp/slices"
)

const (

	// Bind: table
	sqlGetMigrated = `SELECT name FROM %s`

	// Bind: table, name, batch
	sqlInsertMigration = `INSERT INTO %s VALUES (NULL, %#v, %d)`
)

// Migration statuses
// @see Filter.Status
const (
	StatusMigrated    = -1
	StatusNotMigrated = 1
)

type Migrator interface {
	CreateMigration(string) (*Migration, error)
	MigrationExists(string) (bool, error)
	Init() error
	EnsureTable() error
	SetConfig(*config.MigrateConfig)
}

type Migrate struct {
	Config *config.MigrateConfig
	DB     *DB

	// migrated is slice of migration names in DB that are already migrated
	migrated []string
}

// Filter helps to filter all migration files available
type Filter struct {
	//          current
	//     -1  <=  |  => 1
	//  M1  |  M2  |  M3  |  M4  |  M5  |  M6  |
	// -1 migrated (left) , 1 unmigrated ( right )
	Status uint8
}

// NewMigrate craetes new instance of the Migrate (Manager)
// MIgrator also requeres MigrateConfig, to create with config
// see NewMigrateWithConfig
func NewMigrate(db *DB) *Migrate {
	return &Migrate{
		DB: db,
	}
}

// NewMigrateWithConfig creates new instance of Migrate, and
// immediately sets the proper config of MigrateConfig
func NewMigrateWithConfig(db *DB, cfg *config.MigrateConfig) (m *Migrate) {
	m = NewMigrate(db)
	m.Config = cfg
	return
}

// Init prepares migrations table, etc.
func (o *Migrate) Init() (err error) {
	err = o.EnsureTable()
	if err != nil {
		return
	}

	// Load migratied names from the database,
	// and assign it for later use:
	o.migrated, err = o.Migrated()
	if err != nil {
		return err
	}

	return o.Migrate()
}

// CreateMigration accepts name like "create_users_table"
// or "Create Users Table" and creates migration file
func (o *Migrate) CreateMigration(name string) (m *Migration, err error) {
	m = NewMigration(o, name)
	m.Migrate = o

	err = m.Create()
	return
}

// MigrationExists checks if the migration file exists
//  in the migrations directory
func (o *Migrate) MigrationExists(name string) (bool, error) {
	m := NewMigration(o, name)
	return m.Exists()
}

// Migrate executes all migrations that are not already
// executed in up direction
func (o *Migrate) Migrate() (err error) {
	o.DB.Logger.Verbose("Trying to migrate DB...")

	files, err := o.FilteredFiles(Filter{
		Status: StatusNotMigrated,
	})
	if err != nil {
		return err
	}

	name := ""

	o.DB.Logger.Verbose(" |_ Files: ")
	for i := 0; i < len(files); i++ {
		name = path.Base(files[i])
		status := "migrated <-"

		if !slices.Contains(o.migrated, name) {

			status = "not migrated ->"
			m := FromName(o, name)
			err = o.Apply(m)
			if err != nil {
				return err
			}

			err = o.RecordMigrated(name, 1)
			if err != nil {
				return err
			}
		}
		o.DB.Logger.Verbose("   |_ %d. %v - %s", i, name, status)
	}

	return nil
}

func (o *Migrate) Apply(m *Migration) error {
	c, err := m.Read()
	if err != nil {
		return err
	}
	// Execute
	o.DB.Logger.Verbose("SQL to EXEC:")
	o.DB.Logger.Verbose(string(c))
	return o.DB.Adapter.Exec(c)
}

// AbsDir returns absolute directory to the migrations
func (o *Migrate) AbsDir() (string, error) {
	return o.Path(o.Config.Dir)
}

// Path will convert relative path to absolute path
func (o *Migrate) Path(p string) (string, error) {
	abs := filepath.IsAbs(p)
	if abs {
		return p, nil
	}
	mPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return string(mPath + "/" + p), nil
}

// Files returns all migration files available, no matter
// if they are already executed or not
func (o *Migrate) Files() ([]string, error) {
	dirFull, err := o.AbsDir()
	if err != nil {
		return nil, err
	}
	o.DB.Logger.Debug("Migration files from: %s", dirFull)
	files, err := filepath.Glob(path.Join(dirFull, "*.sql"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		o.DB.Logger.Debug(file)
	}

	return files, nil
}

// Returns slice of files filtered by status, etc.
func (o *Migrate) FilteredFiles(filter Filter) ([]string, error) {
	all, err := o.Files()
	if err != nil {
		return nil, err
	}

	// TODO:
	return all, nil
}

// EnsureTable ensures that the migrations table in the DB exists
func (o *Migrate) EnsureTable() error {
	o.DB.Logger.Verbose("Ensuring migrations table exists: %s", o.Config.Table)
	return o.DB.Adapter.InitMigrations(o.Config.Table)
}

// In some cases, we want to update the migrator config, so this func
// provides a way to do it
func (o *Migrate) SetConfig(cfg *config.MigrateConfig) {
	o.Config = cfg
}

// Migrated returns all migrated migrations (string names) from the database
func (o *Migrate) Migrated() (ss []string, err error) {
	ss, err = o.DB.Adapter.Values(sqlGetMigrated, o.Config.Table)
	return ss, err
}

// Migrated returns all migrated migrations (string names) from the database
func (o *Migrate) RecordMigrated(name string, batch uint32) error {
	return o.DB.Adapter.Exec(sqlInsertMigration, o.Config.Table, name, batch)
}
