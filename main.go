package main

import (
	covidtracker "covidtracker/src"
	"os"
)

func main() {
	if err := covidtracker.RunApp(os.Args[1]); err != nil {
		panic(err)
	}
}
