package db

import (
	"encoding/json"
	"io/ioutil"
	"os"

	uno "github.com/glugox/uno/pkg"
	"github.com/glugox/uno/pkg/db"
	"github.com/glugox/uno/pkg/schema"
)

func CRMSeed(db *db.DB) error {
	db.Logger.Info("Seeding CRM...")

	app, err := uno.Instance()
	if err != nil {
		return err
	}

	jsonFile, err := os.Open("./db/seeds/menu.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Parse menu json into Menu struct
	var menu schema.Menu

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &menu)
	if err != nil {
		return err
	}

	err = app.Entity(menu).Save()
	if err != nil {
		return err
	}
	// dp := uno.Entity().New(&User{}).Save()
	/*err := app.Entity(&schema.Menu{}).Insert()
	if err != nil {
		panic(err)
	}*/

	return nil
}
