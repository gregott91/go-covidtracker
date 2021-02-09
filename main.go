package main

import (
	covidtracker "covidtracker/src"
	"os"
)

func main() {
	fileName := "test.json"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	if err := covidtracker.RunApp(fileName); err != nil {
		panic(err)
	}
}
