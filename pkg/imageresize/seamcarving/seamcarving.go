package seamcarving

import (
	"image"
	"math"
)

func ComputePixelEnergy(img image.Image) [][]float64 {

	bounds := img.Bounds()
	top := bounds.Min.Y
	left := bounds.Min.X
	bottom := bounds.Max.Y
	right := bounds.Max.X

	energyData := make([][]float64, (bottom - top))
	for colIndex := range energyData {
		energyData[colIndex] = make([]float64, (right - left))
	}

	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			energy := energyAt(x, y, top, left, bottom, right, &img)
			energyData[y][x] = energy
		}
	}

	return energyData
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
