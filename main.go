package main

import (
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		window := new(app.Window) //creating window
		err := run(window) //calling run
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
