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

// Executes the ImageResize app
func (a ImageResizeApp) Run() error {

	log.Println("Reading image ", a.Filename)

	img, err := utils.OpenImage(a.Filename)
	if err != nil {
		log.Println("Error occurred while opening the image", err)
		return err
	}

	log.Println("Computing pixel energy")

	energyData := seamcarving.ComputePixelEnergy(img)
	highlightedImg, err := utils.RenderGrayImageFromArray(energyData)

	if err != nil {
		log.Println("Error occurred while rendering image from energy-data", err)
		return err
	}

	log.Println("Writing image to file")
	err = utils.WriteImage(a.EnergyHighlightFilename, highlightedImg)

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
