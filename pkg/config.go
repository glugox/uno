package uno

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/glugox/uno/pkg/schema"
	"golang.org/x/exp/slices"
)

const (
	// RootDirFlagFile if file that we place at the root of our project to
	// hold some important data, but also to flag us that it is the root
	RootDirFlagFile = ".uno"
)

const (
	maxDirDepth = 10
)

var (
	levelOneDirs = []string{"pkg", "cmd"}
)

// Config struct holds field names for most common data (DB, etc)
// but also has capacity to hold dynamic config registration
// Config must be serializable/unserializable, so dont' put any structs
// with methods here
type Config struct {
	DB     *DBConfig
	Server *ServerConfig
	App    *AppConfig
}

// NewConfig returns new instance of global configuration struct
func NewConfig() (*Config, error) {

	appCfg, err := NewAppConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App:    appCfg,
		DB:     NewDBConfig(),
		Server: NewServerConfig(),
	}

	// Config should look something like this:
	// {
	// 	"DB": {
	// 	  "Host": "127.0.0.1",
	// 	  "Port": 3306,
	// 	  "User": "root",
	// 	  "Password": "root",
	// 	  "DBName": "go_uno_demo",
	// 	  "Migrate": {
	// 		"Dir": "{App.BasePath}/db/migrations",
	// 		"Table": "migrations"
	// 	  }
	// 	},
	// 	"App": {
	// 	  "BasePath": "~/go/src/github.com/glugox/uno/versions/v1/uno"
	// 	}
	// }
	//
	// We want to replace all tpl vars like e.g. {App.BasePath}

	cfg, err = ReplaceTplVars(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// AppConfig
type AppConfig struct {
	BasePath string
	Verbose  bool
}

// NewAppConfig returns new instance of app configuration struct
func NewAppConfig() (*AppConfig, error) {

	basePath, err := GetBasePath()
	if err != nil {
		uErr := fmt.Errorf("could not find Application's base path. Err: %s", err)
		return nil, uErr
	}

	cfg := &AppConfig{
		BasePath: basePath,
	}

	return cfg, nil
}

// DBConfig
type DBConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DSN      string
	Migrate  *MigrateConfig
}

// MigrateConfig
type MigrateConfig struct {
	Dir   string
	Table string
}

// NewConfig returns new Config instance
// for the current environment
func NewDBConfig() *DBConfig {
	conn := Env("DB_CONNECTION", "sqlite3")

	switch conn {
	case schema.DBAdapterSqlite:
		return NewSqliteDBConfig()
	default:
		return NewMySqlDBConfig()
	}

}

// NewConfig returns new Config instance
// for the current environment
func NewMySqlDBConfig() *DBConfig {
	return &DBConfig{
		Name:     DBAdapterMySql,
		Host:     Env("DB_HOST", "127.0.0.1"),       // 127.0.0.1
		Port:     Env("DB_PORT", "3306"),            // 3306
		User:     Env("DB_USERNAME", "root"),        // root
		Password: Env("DB_PASSWORD", "root"),        // root
		DBName:   Env("DB_DATABASE", "go_uno_demo"), // go_uno_demo
		Migrate:  NewMigrateConfig(),
	}
}

// NewConfig returns new Config instance
// for the current environment
func NewSqliteDBConfig() *DBConfig {
	return &DBConfig{
		Name:    DBAdapterSqlite,
		DSN:     Env("DB_DSN", "./uno.db"),
		Migrate: NewMigrateConfig(),
	}
}

// NewMigrateConfig returns migration config
// based on env settings
func NewMigrateConfig() *MigrateConfig {
	return &MigrateConfig{

		// The default value of "migrations" will result in
		// /db/migrations at the root of the project
		Dir: "{App.BasePath}/db/migrations",

		// Database table name
		Table: "migrations",
	}
}

// ServerConfig
type ServerConfig struct {
	Address string
}

// NewServerConfig creates new instance of ServerConfig
// with default values
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Address: ":9090",
	}
}

// GetBasePath returns applications base path
func GetBasePath() (base string, err error) {
	curr, osErr := os.Getwd()
	currOrig := curr
	if osErr != nil {
		return "", osErr
	}

	// Here we want to check if we are already at the base path
	dirs, err := os.ReadDir(curr)
	if err != nil {
		return "", err
	}

	// Check if at our current directory we have a file that marks
	// the root directory of our app
	for _, dir := range dirs {
		if dir.Name() == RootDirFlagFile {
			// We are already at the root
			return curr, nil
		}
	}

	base = ""
	err = nil
	for i := 0; i < maxDirDepth; i++ {
		base = path.Base(curr)
		if slices.Contains(levelOneDirs, base) {
			base = path.Dir(curr)
			err = nil
			return
		}
		curr = path.Dir(curr)
	}

	// At this point, return current working dir
	return currOrig, err
}

// ToString implements config.Node interface
func (c *DBConfig) ToString() string {
	if len(c.DSN) > 0 {
		return c.DSN
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

// ReplaceTplVars replaces tpl vars like {App.BasePath} with actuall values
func ReplaceTplVars(cfg *Config) (*Config, error) {

	logger := DefaultLogFactory().NewLogger()
	var newCfg Config
	bJson, err := json.Marshal(cfg)
	if err != nil {
		return nil, logger.Error("could not marshal main config")
	}

	strJson := string(bJson)

	// Regexp to find all tpl vars
	re, err := regexp.Compile(`\{([a-zA-Z_\.]+)\}`)
	if err != nil {
		return nil, logger.Error("could not parse main config")
	}

	valCfg := reflect.ValueOf(cfg).Elem()

	// Replacer of single tpl var e.g. "{App.BasePath}"
	var replacer = func(s string) string {
		s = strings.Trim(s, "{}")
		curr := valCfg
		parts := strings.Split(s, ".")
		for _, part := range parts {
			curr = curr.FieldByName(part)
			if curr.Kind() == reflect.Pointer {
				curr = curr.Elem()
			}
		}

		return curr.String()
	}

	// Replace all tpl vars
	strJson = re.ReplaceAllStringFunc(strJson, func(s string) string {
		return replacer(s)
	})

	logger.Debug("Config: %v", strJson)

	err = json.Unmarshal([]byte(strJson), &newCfg)
	if err != nil {
		return nil, err
	}

	return &newCfg, nil
}
