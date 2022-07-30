package main

import (
	"fmt"

	uno "github.com/glugox/uno/pkg"
)

func main() {

	uno, err := uno.New()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Basic Example: %s \n", uno.Config.Server.Address)
}
