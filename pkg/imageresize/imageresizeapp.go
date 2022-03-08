package imageresize

import (
	"flag"
	"fmt"
	"log"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/seamcarving"
	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/utils"
)

type ImageResizeApp struct {
	Filename          string
	OutputFilename    string
	ResizeAdjustWidth int
}

// Creates a new instance of the ImageResizeApp after
// validating the argument list
func New(args []string) (*ImageResizeApp, error) {

	resizeAdjustWidth := flag.Int("width", 10, "No. of pixels of width to remove from image")
	newFileNameFlag := flag.String("write", "", "Image path to write output to")

	flag.Parse()
	remainingArgs := flag.Args()

	if len(remainingArgs) != 1 {
		return nil, fmt.Errorf("not enough arguments")
	}

	imageFileName := remainingArgs[0]

	newFileName := *newFileNameFlag
	if *newFileNameFlag == "" {
		newFileName = fmt.Sprintf("%s-resized.jpeg", imageFileName)
	}

	app := &ImageResizeApp{
		Filename:          imageFileName,
		OutputFilename:    newFileName,
		ResizeAdjustWidth: *resizeAdjustWidth,
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
	log.Println("Resizing width by ", a.ResizeAdjustWidth)
	for i := 0; i < a.ResizeAdjustWidth; i++ {

		log.Print(i + 1)
		energyData := seamcarving.ComputePixelEnergy(img)
		seamCoords := seamcarving.ComputeVerticalSeam(energyData)
		carvedImage := utils.RemoveVerticalSeamFromImage(&img, seamCoords)

		img = carvedImage
	}

	log.Println("Writing output to ", a.OutputFilename)
	err = utils.WriteRGBAImage(a.OutputFilename, &img)
	if err != nil {
		log.Println("Error occurred while writing image to the filesystem", err)
		return err
	}

	log.Println("Complete")

	return nil
}
