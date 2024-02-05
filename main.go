package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string = "problems.csv"
	var language string = "en"
	var shuffle bool = false
    var timer int = 30
	if len(os.Args) >= 2 {
        if os.Args[1] == "help" {
            fmt.Println(GetText(language, "help"))
            return
        }
	}
	InitGame(filename, timer, language, shuffle)
}

