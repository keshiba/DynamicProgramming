package imageresize

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/seamcarving"
)

type ImageResizeApp struct {
	Filename                string
	EnergyHighlightFilename string
}

func (a ImageResizeApp) Run() error {

	img, err := openImage(a.Filename)
	if err != nil {
		return err
	}

	energyData := seamcarving.ComputePixelEnergy(img)
	highlightedImg, err := seamcarving.RenderImageEnergy(energyData)
	if err != nil {
		return err
	}

	err = writeImage(a.EnergyHighlightFilename, highlightedImg)

	return err
}

func openImage(fileName string) (image.Image, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func writeImage(outFileName string, img *image.Gray16) error {

	newF, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	defer newF.Close()

	jpegOptions := jpeg.Options{
		Quality: 90,
	}

	err = jpeg.Encode(newF, img, &jpegOptions)

	return err
}
