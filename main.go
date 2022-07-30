package main

import (
	"log"

	"github.com/glugox/uno/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	cmd.Execute()
}
