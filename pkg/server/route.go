package server

import (
	"net/http"
	"regexp"
)

// Route is an entity to hold matching data between route
// and its hanlder.
type Route struct {
	Method  string
	Regex   *regexp.Regexp
	Handler http.HandlerFunc
}

// NewRoute creates instance of new Route
func NewRoute(method string, pattern string, handler http.HandlerFunc) Route {
	return Route{Method: method, Regex: regexp.MustCompile("^" + pattern + "$"), Handler: handler}
}
