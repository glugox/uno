package db

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/glugox/uno/pkg/schema"
	"github.com/glugox/uno/pkg/utils"
)

// Migration is mirror to single migration file
type Migration struct {

	// Pointer to the Migrate (Manager)
	Migrate *Migrate

	// Id mirror in database
	Id schema.ObjectId

	// Name is migration file name (no dir) and also database name
	Name string

	// While Name can be changed with adding timestamp and extension,
	// OrigName is the name we used initially to create the migration
	OrigName string
}

// NewMigration creates new migration struct without ceateing any file
// in the system. It takes migration manager (mm) as the argument,
// so it has access to the config. To create migration use:
// Migrate.CreateMigration
func NewMigration(mm *Migrate, name string) *Migration {

	// Keep track of the original migration name separately
	origName := name

	// Year: 2006, Month: 01, Days: 02, Hours: 15, Minutes: 04, Seconds: 05
	ts := time.Now().Format("20060102150405")

	ext := ".sql"

	// Check if name ends up with en extension. In that case remove it
	// from the name and assign it to the ext var so we append it again
	if newExt := filepath.Ext(name); newExt != "" {
		ext = newExt // Ext retruns the dot also
	}

	// If at this point we already have extension in the name
	// we need to remove it because we will add it again
	name = strings.Trim(name, ext)

	return &Migration{
		Id:       schema.NewObjectId(),
		Name:     fmt.Sprintf("%s_%s%s", ts, utils.ToSnakeCase(name), ext),
		OrigName: origName,
		Migrate:  mm,
	}
}

// FromName retruns new Migration instance based on the (file) name.
// This is reversing of the situation when we want to get file name
// based on some string like "Create Users Table". Here we want to
// get a Migration from e.g. "200000000000_create_users_table.sql"
func FromName(mm *Migrate, name string) *Migration {
	m := &Migration{
		Id:       schema.NewObjectId(),
		Name:     name,
		OrigName: name,
		Migrate:  mm,
	}

	return m
}

// CreateAt creates (writes) migration file at specific directory
func (m *Migration) Create() (err error) {

	t := m.ExtractTableName("<table_name>")
	sql := fmt.Sprintf(`CREATE TABLE %s
	(
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		created_at   timestamp    null,
		updated_at   timestamp    null
	)`, t)

	m.Write(sql)

	return nil
}

// Remove deletes the migration file from thy system
func (o *Migration) Remove() (err error) {
	fPath, err := o.FullPath()
	if err != nil {
		return err
	}

	err = os.Remove(fPath)
	return
}

// ExtractTableName tries to extract table name from the migration name.
// e.g. '20220712121033_create_users_table.sql' will extract 'users'
func (m *Migration) ExtractTableName(alt string) (t string) {
	re, _ := regexp.Compile("create_([a-z]+)_table")

	for _, aT := range re.FindAllStringSubmatch(m.Name, -1) {
		t = aT[1] // 0 should be full match
		return
	}

	if len(t) == 0 {
		return alt
	}
	return
}

// Write sets the content of our new migration file
func (o *Migration) Write(sql string) (err error) {

	path, err := o.FullPath()

	// Migrate's only boss is DB, so we can borrow logger
	o.Migrate.DB.Logger.Info("Creating migration : %s", path)
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	// Write our SQL to the migration file
	_, err = f.WriteString(sql)
	return
}

// Read gets the contents from our migration file
func (o *Migration) Read() (c string, err error) {

	path, err := o.FullPath()
	if err != nil {
		return
	}

	bc, err := os.ReadFile(path)
	if err != nil {
		return
	}
	c = string(bc)

	return
}

// FullPath returns the absolute path of the migration file
func (o *Migration) FullPath() (string, error) {
	dir, err := o.Migrate.Path(o.Migrate.Config.Dir)
	if err != nil {
		return "", err
	}
	return string(dir + "/" + o.Name), nil
}

// Exists returns true if the migration file exists
func (o *Migration) Exists() (bool, error) {

	// Get full path of migration
	path, err := o.FullPath()
	if err != nil {
		return false, err
	}

	// Check file exists in the system
	_, err = os.Stat(path)
	if err != nil {
		// No need to pass back PathError , just false
		return false, nil
	}
	return true, nil
}
