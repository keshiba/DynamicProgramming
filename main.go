package main

import (
	"fmt"

	"github.com/keshiba/ll-dynamicprogramming/pkg/algorithm/flowerbox"
)

func main() {

	box := []int{}

	maxHeight := flowerbox.CalculateMaxFlowerBoxHeight(box)

	prompt := fmt.Sprintf("Max Height = %d", maxHeight)
	fmt.Println(prompt)
}
