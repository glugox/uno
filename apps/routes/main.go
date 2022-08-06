package main

import (
	"net/http"

	uno "github.com/glugox/uno/pkg"
	"github.com/glugox/uno/pkg/schema"
	"github.com/glugox/uno/pkg/server"
)

func main() {

	uno, err := uno.Instance()
	if err != nil {
		panic(err)
	}

	uno.Init()
	uno.RegisterRoutes(Routes())
	uno.Run()
}

func Routes() []*server.Route {
	var routes = []*server.Route{
		/*
			uno.Get("/menu/([^/]+)/items/([0-9]+)/update", MenusItemsUpdate),
		*/

		uno.Get("/menus", MenusGet),
		uno.Get("/menus/([^/]+)", MenusItemGet),
		uno.Get("/menus-items", MenusItemsGet),
	}
	return routes
}

// MenusGet returns list of all our menus
func MenusGet(w http.ResponseWriter, r *http.Request) {
	uno, err := uno.Instance()
	if err != nil {
		panic(err)
	}

	// dp := uno.Entity().New(&User{}).Save()
	dp, err := uno.Entity(&schema.Menu{}).All()
	if err != nil {
		panic(err)
	}

	// Convert our data provider to json
	res, err := dp.Marshal()
	if err != nil {
		panic(err)
	}

	w.Write(res)
}

// MenusGet returns one menu ( not menu item as child of menu )
// e.g. Admin Menu, Settings Menu
func MenusItemGet(w http.ResponseWriter, r *http.Request) {
	menuName := server.RequestParam(r, 0)
	uno, err := uno.Instance()
	if err != nil {
		panic(err)
	}

	// dp := uno.Entity().New(&User{}).Save()
	dp, err := uno.Entity(&schema.Menu{}).Where("id", menuName).First()
	if err != nil {
		panic(err)
	}

	// Convert our data provider to json
	res, err := schema.MarshalOnlyModelDefeult(dp)
	if err != nil {
		panic(err)
	}

	w.Write(res)
}

func MenusItemsGet(w http.ResponseWriter, r *http.Request) {
	uno, err := uno.Instance()
	if err != nil {
		panic(err)
	}

	// dp := uno.Entity().New(&User{}).Save()
	dp, err := uno.Entity(&schema.MenuItem{}).All()
	if err != nil {
		panic(err)
	}

	// Convert our data provider to json
	res, err := dp.Marshal()
	if err != nil {
		panic(err)
	}

	w.Write(res)
}
