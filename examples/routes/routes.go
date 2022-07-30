package main

import (
	"net/http"

	uno "github.com/glugox/uno/pkg"
)

func Routes() []uno.Route {
	var routes = []uno.Route{
		uno.NewRoute("GET", "/", UsersGet),
		uno.NewRoute("GET", "/contact", UsersGet),
		uno.NewRoute("GET", "/api/widgets", UsersGet),
		uno.NewRoute("POST", "/api/widgets", UsersGet),
		uno.NewRoute("POST", "/api/widgets/([^/]+)", UsersGet),
		uno.NewRoute("POST", "/api/widgets/([^/]+)/parts", UsersGet),
		uno.NewRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", UsersGet),
		uno.NewRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/delete", UsersGet),
		uno.NewRoute("GET", "/([^/]+)", UsersGet),
		uno.NewRoute("GET", "/([^/]+)/admin", UsersGet),
		uno.NewRoute("POST", "/([^/]+)/image", UsersGet),
	}
	return routes
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

// Users

// Handles POST /api/widgets/([^/]+)/parts/([0-9]+)/update
func UsersGet(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "UsersGet %s %d\n", r.String(0), r.Int(0))
	// Get all Users from our DB
	// ctx.Entity(&User, &Roles).All().asJSON()
}

// UsersIndex GET: /users/1
func UsersGetItem(ctx uno.Context) {
	// Get one User from DB
	// ctx.Entity(&User).First().asJSON()
}

// UsersIndex POST: /users/
func UsersPost() {
	// ctx.Entity(&User).First().asJSON()
}

// UsersIndex POST: /users/1
func UsersEditItem() {

}

// UsersIndex DELETE: /users/1
func UsersDeleteItem() {

}

// Menu

// MenuIndex GET: /menu/
func MenuGet(ctx uno.Context) {
	// Get all Menu from our DB
	// ctx.Entity(&User, &Roles).All().asJSON()
}

// MenuIndex GET: /menu/1
func MenuGetItem(ctx uno.Context) {
	// Get one User from DB
	// ctx.Entity(&User).First().asJSON()
}

// MenuIndex POST: /menu/
func MenuPost() {
	// ctx.Entity(&User).First().asJSON()
}

// MenuIndex POST: /menu/1
func MenuEditItem() {

}

// MenuIndex DELETE: /menu/1
func MenuDeleteItem() {

}
