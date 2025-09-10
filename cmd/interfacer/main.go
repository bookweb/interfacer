package main

import (
	"log"
	"os"

	"github.com/bookweb/interfacer/cmd/interfacer/internal/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
