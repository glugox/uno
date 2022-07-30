package main

import (
	"fmt"
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
	fmt.Println(uno.Config.App.Key)
	uno.Run()
}

func Routes() []*server.Route {
	var routes = []*server.Route{
		/*
			uno.Get("/menu/([^/]+)/items/([0-9]+)/update", MenusItemsUpdate),
		*/

		uno.Get("/menu", MenusGet),
		uno.Get("/menu-items", MenusItemsGet),
	}
	return routes
}

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
