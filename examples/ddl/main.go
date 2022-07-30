package main

import (
	uno "github.com/glugox/uno/pkg"
)

func main() {
	ddl := uno.NewDDL()
	err := ddl.Configure(&uno.User{}, &uno.Role{})
	if err != nil {
		panic(err)
	}
	schema := ddl.Read()
	_, err = schema.Dump()
	if err != nil {
		panic(err)
	}

	//fmt.Print(json)

}
