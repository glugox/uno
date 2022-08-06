// CRM application
package main

import (
	"github.com/glugox/uno/apps/crm/db"
	"github.com/glugox/uno/apps/crm/models"
	"github.com/glugox/uno/apps/crm/routes"
	uno "github.com/glugox/uno/pkg"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	uno, err := uno.Instance()
	if err != nil {
		panic(err)
	}

	// RegisterModels must come before uno.Init()
	// as Init will initialize database
	uno.RegisterModels(models.AllModels())
	uno.RegisterSeeder(db.CRMSeed)
	uno.RegisterRoutes(routes.AllRoutes())

	err = uno.Init()
	if err != nil {
		panic(err)
	}

	uno.Run()
}
