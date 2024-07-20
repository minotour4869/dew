package main

import (
	"log"

	"github.com/minotour4869/dew/cmd"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	cmd.Execute()
}
