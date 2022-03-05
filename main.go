package main

import (
	"log"
	"os"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize"
)

func main() {

	app, err := imageresize.New(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
