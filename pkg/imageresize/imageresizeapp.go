package imageresize

import (
	"fmt"
	"log"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/seamcarving"
	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/utils"
)

type ImageResizeApp struct {
	Filename                string
	EnergyHighlightFilename string
}

// Creates a new instance of the ImageResizeApp after
// validating the argument list
func New(args []string) (*ImageResizeApp, error) {

	if len(args) != 2 {
		printUsage(args[0])
		return nil, fmt.Errorf("not enough arguments")
	}

	imageFileName := args[1]
	newFileName := fmt.Sprintf("%s-energy.jpeg", imageFileName)

	app := &ImageResizeApp{
		Filename:                imageFileName,
		EnergyHighlightFilename: newFileName,
	}

	return app, nil
}

// Runs the ImageResizeApp which computes the energy value
// of each pixel in an image and performs content-aware resizing
// by carving out seams of pixels from the image with low energy
func (a ImageResizeApp) Run() error {

	log.Println("Reading image ", a.Filename)

	img, err := utils.OpenImage(a.Filename)
	if err != nil {
		log.Println("Error occurred while opening the image", err)
		return err
	}

	log.Println("Low energy seam carving in progress")
	for i := 0; i < 400; i++ {

		log.Print(i + 1)
		energyData := seamcarving.ComputePixelEnergy(img)
		seamCoords := seamcarving.ComputeVerticalSeam(energyData)
		carvedImage := utils.RemoveVerticalSeamFromImage(&img, seamCoords)

		img = carvedImage
	}

	carvedImgFileName := fmt.Sprintf("%s-%s", a.Filename, "carved.jpg")
	err = utils.WriteRGBAImage(carvedImgFileName, &img)
	if err != nil {
		log.Println("Error occurred while writing image to the filesystem", err)
		return err
	}

	log.Println("Complete")

	return nil
}

func printUsage(binaryName string) {

	fmt.Printf("Usage: ./%s <input-file-name>\n", binaryName)
}
