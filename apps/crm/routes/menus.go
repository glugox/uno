package routes

import (
	"github.com/glugox/uno/apps/crm/controllers"
	uno "github.com/glugox/uno/pkg"
	"github.com/glugox/uno/pkg/server"
)

func MenusRoutes() []*server.Route {
	var routes = []*server.Route{
		/*
			uno.Get("/menu/([^/]+)/items/([0-9]+)/update", MenusItemsUpdate),
		*/

		uno.Get("/menus", controllers.MenusGet),
		uno.Get("/menus/([^/]+)", controllers.MenusItemGet),
		uno.Get("/menus-items", controllers.MenusItemsGet),
	}
	return routes
}
