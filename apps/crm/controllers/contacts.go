package controllers

import (
	"net/http"

	"github.com/glugox/uno/apps/crm/models"
	uno "github.com/glugox/uno/pkg"
)

// ContactsGet returns list of all our contacts
func ContactsGet(w http.ResponseWriter, r *http.Request) {
	uno, err := uno.Instance()
	if err != nil {
		panic(err)
	}

	dp, err := uno.Entity(&models.Contact{}).All()
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
