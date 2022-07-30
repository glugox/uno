package main

import (
	uno "github.com/glugox/uno/pkg"
	schema "github.com/glugox/uno/pkg/schema"
)

func main() {

	app, _ := uno.New()
	app.Init()

	// dp := uno.Entity().New(&User{}).Save()
	dp, err := app.Entity(&schema.Menu{}).All()

	if err != nil {
		panic(err)
	}
	dp.Print()
}
