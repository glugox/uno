package uno

import (
	"net/http"
	"regexp"
	"sync"

	"github.com/glugox/uno/pkg/config"
	"github.com/glugox/uno/pkg/db"
	"github.com/glugox/uno/pkg/log"
	"github.com/glugox/uno/pkg/schema"
	"github.com/glugox/uno/pkg/server"
)

var once sync.Once
var instance *Uno

// UnoApi is an Uno Application for creating APIs
type Uno struct {
	Server *server.Server
	DB     *db.DB
	Config *config.Config
	Logger log.Logger
}

// Init initializes DB, etc
func (o *Uno) Init() error {
	o.Logger.Info("Uno init...")
	return o.DB.Init()
}

// New creates new Api applications
func New() (u *Uno, err error) {
	// Global configuration
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	return WithConfig(cfg)
}

// Instance retuns singleton instance of our app
func Instance() (u *Uno, err error) {
	once.Do(func() {
		instance, err = New()
	})

	return instance, err
}

// WithConfig Creates new Uno app with already defined config
func WithConfig(cfg *config.Config) (u *Uno, err error) {
	// Global configuration
	if err != nil {
		return nil, err
	}

	// App server
	srv := server.NewServer(cfg.Server.Address)

	//dbo := db.NewDB()
	dbo, err := db.DBWithConfig(cfg.DB)
	if err != nil {
		return nil, err
	}

	// Main app
	u = &Uno{
		Config: cfg,
		Server: srv,
		DB:     dbo,
		Logger: log.DefaultLogFactory().NewLogger(),
	}

	// Add app reference to the Children
	u.Server.Init()

	return u, nil
}

func (u *Uno) RegisterRoutes(routes []*server.Route) {
	u.Server.RegisterRoutes(routes)
}

// Run executes the applications
func (u *Uno) Run() {
	u.Server.Run()
}

// Entity returns new database session
func (u *Uno) Entity(model schema.Model) *db.Session {
	return db.NewSession(u.DB, model)
}

func Get(pattern string, handler http.HandlerFunc) *server.Route {
	return &server.Route{Method: "GET", Regex: regexp.MustCompile("^" + pattern + "$"), Handler: handler}
}

func Post(pattern string, handler http.HandlerFunc) *server.Route {
	return &server.Route{Method: "POST", Regex: regexp.MustCompile("^" + pattern + "$"), Handler: handler}
}
