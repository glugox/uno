package routes

import "github.com/glugox/uno/pkg/server"

func AllRoutes() []*server.Route {
	//return MenusRoutes()
	return append(MenusRoutes(), ContactsRoutes()...)
}
