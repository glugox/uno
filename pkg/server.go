package uno

import "net/http"

// Server is internal API server for handling requests
type Server struct {
	Address    string
	App        *Uno
	httpServer *http.Server
	Mux        *http.ServeMux
}
