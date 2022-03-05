package utils

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

// RenderGrayImageFromArray converts a 2D array of values into
// a Grayscale image
func RenderGrayImageFromArray(imgArray [][]float64) (*image.Gray16, error) {

	if len(imgArray) == 0 {
		return nil, errors.New("energy data is empty")
	}

	height := len(imgArray)
	width := len(imgArray[0])

	newRect := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}
	newImg := image.NewGray16(newRect)

	for y, rowData := range imgArray {
		for x, pixelValue := range rowData {

			newColor := color.Gray16{
				Y: uint16(pixelValue),
			}

			newImg.SetGray16(x, y, newColor)
		}
	}

	return newImg, nil
}

// OpenImage reads an image from the filesystem
func OpenImage(fileName string) (image.Image, error) {

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

// WriteImage writes a 16bit Grayscale image (image.Gray16)
// to the filesystem
func WriteImage(outFileName string, img *image.Gray16) error {

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
