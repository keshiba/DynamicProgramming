package seamcarving

import (
	"image"
	"math"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/types"
)

type EnergyData struct {
	MinEnergy    float64
	MinParentCol int
}

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
			energyData[y][x] = energyAt(x, y, top, left, bottom, right, &img)
		}
	}

	return energyData
}

func ComputeVerticalSeam(energyData [][]float64) []types.Coordinate {

	height := len(energyData)
	width := len(energyData[0])

	minEnergyGrid := make([][]EnergyData, height)
	for rowIndex := range minEnergyGrid {
		minEnergyGrid[rowIndex] = make([]EnergyData, width)
		for colIndex := range minEnergyGrid[rowIndex] {
			minEnergyGrid[rowIndex][colIndex] = EnergyData{0, 0}
		}
	}

	for colIndex := 0; colIndex < width; colIndex++ {
		minEnergyGrid[0][colIndex] =
			EnergyData{energyData[0][colIndex], 0}
	}

	for rowIndex := 1; rowIndex < height; rowIndex++ {
		for colIndex := 0; colIndex < width; colIndex++ {
			prevRowMinCol := colIndex - 1
			if prevRowMinCol < 0 {
				prevRowMinCol = 0
			}

			prevRowMaxCol := colIndex + 1
			if prevRowMaxCol > (width - 1) {
				prevRowMaxCol = width - 1
			}

			minParentCol := minPosOfEnergyData(minEnergyGrid[rowIndex-1], prevRowMinCol, prevRowMaxCol+1)

			minEnergyGrid[rowIndex][colIndex] =
				EnergyData{
					energyData[rowIndex][colIndex] + minEnergyGrid[rowIndex-1][minParentCol].MinEnergy,
					minParentCol,
				}
		}
	}

	minCol := minPosOfEnergyData(minEnergyGrid[height-1], 0, width)

	minEnergySeamCoordinates := make([]types.Coordinate, height)

	colIndex := minCol
	for rowIndex := height - 1; rowIndex >= 0; rowIndex-- {
		energyValue := minEnergyGrid[rowIndex][colIndex]
		minEnergySeamCoordinates[rowIndex] = types.Coordinate{Row: rowIndex, Col: colIndex}

		colIndex = energyValue.MinParentCol
	}

	return minEnergySeamCoordinates
}

func minPosOfEnergyData(a []EnergyData, start, end int) int {

	minValue := a[start].MinEnergy
	minValuePos := start

	for i := start + 1; i < end; i++ {

		if a[i].MinEnergy < minValue {
			minValue = a[i].MinEnergy
			minValuePos = i
		}
	}

	return minValuePos
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
