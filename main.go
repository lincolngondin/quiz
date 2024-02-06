package main

import (
	"log"
	"os"
)

func main() {
	opts, helpCommand, err := ProcessArguments(os.Args)
	if helpCommand {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	InitGame(opts)
}
