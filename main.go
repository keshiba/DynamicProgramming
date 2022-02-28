package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/seamcarving"
)

func main() {

	imageFileName := `D:\Dev\golang\ll-dynamicprog\resource\img\surfer.jpg`

	fmt.Printf("Opening image file at %s\n", imageFileName)

	f, err := os.Open(imageFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	highlightedImg := seamcarving.HighlightImageEnergy(img)

	newFileName := fmt.Sprintf("%s-energy.jpeg", imageFileName)
	newF, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer newF.Close()

	jpegOptions := jpeg.Options{
		Quality: 90,
	}

	fmt.Println("Writing image")
	err = jpeg.Encode(newF, highlightedImg, &jpegOptions)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Complete")
}
