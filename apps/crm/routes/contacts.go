package routes

import (
	"github.com/glugox/uno/apps/crm/controllers"
	uno "github.com/glugox/uno/pkg"
	"github.com/glugox/uno/pkg/server"
)

func ContactsRoutes() []*server.Route {
	var routes = []*server.Route{
		uno.Get("/contacts", controllers.ContactsGet),
	}
	return routes
}
