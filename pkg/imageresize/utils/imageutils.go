package utils

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/types"
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

func WriteRGBAImage(outFileName string, img *image.Image) error {

	newF, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	defer newF.Close()

	jpegOptions := jpeg.Options{
		Quality: 90,
	}

	err = jpeg.Encode(newF, *img, &jpegOptions)

	return err
}

func HighlightSeam(img *image.Image, coords []types.Coordinate) *image.RGBA64 {

	bounds := (*img).Bounds()
	top := bounds.Min.Y
	left := bounds.Min.X
	bottom := bounds.Max.Y
	right := bounds.Max.X
	width := right - left
	height := bottom - top

	newRect := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}
	newImg := image.NewRGBA64(newRect)

	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			newImg.Set(x, y, (*img).At(x, y))
		}
	}

	for rowIndex := range coords {
		rowCoord := coords[rowIndex].Row
		colCoord := coords[rowIndex].Col

		col := color.RGBA64{R: math.MaxUint16, G: 0, B: 0, A: color.Opaque.A}
		newImg.Set(colCoord, rowCoord, col)

	}

	return newImg
}

// RemoveVerticalSeamFromImage removes a seam of vertically connected pixels
// and returns an image that is 1 pixel smaller in width
func RemoveVerticalSeamFromImage(img *image.Image, coords []types.Coordinate) *image.RGBA64 {

	bounds := (*img).Bounds()
	top := bounds.Min.Y
	left := bounds.Min.X
	bottom := bounds.Max.Y
	right := bounds.Max.X
	width := right - left
	height := bottom - top

	newRect := image.Rectangle{image.Point{0, 0}, image.Pt(width-1, height)}
	newImg := image.NewRGBA64(newRect)

	seamCoordMap := make(map[types.Coordinate]bool)
	for _, coord := range coords {
		seamCoordMap[coord] = true
	}

	for y := top; y < bottom; y++ {
		xAdjust := 0

		for x := left; x < right; x++ {
			adjustedX := x - xAdjust

			coord := types.Coordinate{Row: y, Col: x}
			if _, exists := seamCoordMap[coord]; exists {
				xAdjust = 1
				continue
			}

			color := (*img).At(x, y)
			newImg.Set(adjustedX, y, color)
		}
	}

	return newImg
}

func HighlightPixel(img *image.Gray16, row, col int) error {

	bounds := img.Bounds()
	top := bounds.Min.Y
	left := bounds.Min.X
	bottom := bounds.Max.Y
	right := bounds.Max.X

	if row < top || row > (bottom-1) {
		return errors.New("argument row not within limits")
	}

	if col < left || col > (right-1) {
		return errors.New("argument col not within limits")
	}

	highlightPixelValue := color.Gray16{Y: math.MaxUint16}

	if row > top {
		img.SetGray16(row-1, col, highlightPixelValue)

		if col > left {
			img.SetGray16(row-1, col-1, highlightPixelValue)
		}

		if col < (right - 1) {
			img.SetGray16(row-1, col+1, highlightPixelValue)
		}
	}

	if row < (bottom - 1) {
		img.SetGray16(row+1, col, highlightPixelValue)

		if col > left {
			img.SetGray16(row+1, col-1, highlightPixelValue)
		}

		if col < (right - 1) {
			img.SetGray16(row+1, col+1, highlightPixelValue)
		}
	}

	if col > left {
		img.SetGray16(row, col-1, highlightPixelValue)
	}

	if col < (right - 1) {
		img.SetGray16(row, col+1, highlightPixelValue)
	}

	return nil
}
