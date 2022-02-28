package seamcarving

import (
	"image"
	"image/color"
	"math"
)

func HighlightImageEnergy(img image.Image) *image.Gray16 {

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	top := bounds.Min.Y
	left := bounds.Min.X
	bottom := bounds.Max.Y
	right := bounds.Max.X

	newRect := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}
	newImg := image.NewGray16(newRect)

	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {

			energy := energyAt(x, y, top, left, bottom, right, &img)

			newColor := color.Gray16{
				Y: uint16(energy),
			}

			newImg.SetGray16(x, y, newColor)
		}
	}

	return newImg
}

func energyAt(x, y, top, left, bottom, right int, img *image.Image) float64 {

	xLeft := x - 1
	if xLeft < left {
		xLeft = left
	}

	xRight := x + 1
	if xRight > (right - 1) {
		xRight = right - 1
	}

	yUp := y - 1
	if yUp < top {
		yUp = top
	}

	yDown := y + 1
	if yDown > (bottom - 1) {
		yDown = bottom - 1
	}

	dx := computePixelDifference(xLeft, y, xRight, y, img)
	dy := computePixelDifference(x, yUp, x, yDown, img)

	energy := dx + dy

	return energy
}

func computePixelDifference(x1, y1, x2, y2 int, img *image.Image) float64 {

	lR, lG, lB, _ := (*img).At(x1, y1).RGBA()
	rR, rG, rB, _ := (*img).At(x2, y2).RGBA()

	dRed := lR - rR
	dGreen := lG - rG
	dBlue := lB - rB
	dColor := (dRed * dRed) + (dGreen * dGreen) + (dBlue * dBlue)

	return math.Sqrt(float64(dColor))
}
