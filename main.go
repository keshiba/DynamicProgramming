package main

import (
	"fmt"
	"os"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize"
)

func main() {

	if len(os.Args) != 2 {
		printUsage(os.Args[0])
		return
	}

	imageFileName := os.Args[1]
	newFileName := fmt.Sprintf("%s-energy.jpeg", imageFileName)

	app := imageresize.ImageResizeApp{
		Filename:                imageFileName,
		EnergyHighlightFilename: newFileName,
	}

	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func printUsage(binaryName string) {

	fmt.Printf("Usage: ./%s <input-file-name>\n", binaryName)
}
