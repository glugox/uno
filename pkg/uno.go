package uno

import (
	"github.com/glugox/uno/pkg/config"
	"github.com/glugox/uno/pkg/db"
	"github.com/glugox/uno/pkg/log"
)

// UnoApi is an Uno Application for creating APIs
type Uno struct {
	Server *Server
	DB     *db.DB
	Config *config.Config
	Logger log.Logger
}
