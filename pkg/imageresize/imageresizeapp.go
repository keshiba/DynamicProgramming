package imageresize

import (
	"log"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/seamcarving"
	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/utils"
)

type ImageResizeApp struct {
	Filename                string
	EnergyHighlightFilename string
}

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
