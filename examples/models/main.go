package main

import (
	uno "github.com/glugox/uno/pkg"
)

func main() {

	app, _ := uno.New()
	app.Init()

	// dp := uno.Entity().New(&User{}).Save()

	dp, err := app.Entity(&uno.Menu{}).All()

	if err != nil {
		panic(err)
	}
	dp.Print()
}
