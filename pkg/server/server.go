package server

import (
	"net/http"

	log "github.com/glugox/uno/pkg/log"
)

// IServer interface
type IServer interface {
	GetMux() *http.ServeMux
}

// Server is internal API server for handling requests
type Server struct {
	Address    string
	Logger     log.Logger
	httpServer *http.Server
	Router     *Router
}

// Creates new server instance on a given http address
func NewServer(address string) *Server {

	srv := &Server{
		Address: address,
		httpServer: &http.Server{
			Addr: address,
		},
		Logger: log.DefaultLogFactory().NewLogger(),
		Router: NewRouter(),
	}

	return srv
}

// Init initialises all necessary things for
// the server to run
func (s *Server) Init() {
	s.Logger = log.DefaultLogFactory().NewLogger()
	s.HandleRoutes()

}

func (s *Server) RegisterRoutes(routes []*Route) {
	s.Router.RegisterRoutes(routes)
}

// HandleRoutes set up all the routes
func (s *Server) HandleRoutes() {
	s.httpServer.Handler = s.Router
}

// Run HTTP Server
func (s *Server) Run() error {
	s.Logger.Success("Uno API Running at %s", s.Address)
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}
