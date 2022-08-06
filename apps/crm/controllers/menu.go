package controllers

import (
	"net/http"

	uno "github.com/glugox/uno/pkg"
	"github.com/glugox/uno/pkg/schema"
	"github.com/glugox/uno/pkg/server"
)

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
	dp, err := uno.Entity(&schema.Menu{}).Where("Id", menuName).First()
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
