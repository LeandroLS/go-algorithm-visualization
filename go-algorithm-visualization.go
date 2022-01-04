package main

import (
	"fmt"

	"github.com/oakmound/oak/v3/collision"
)

// Collision labels
const (
	// The only collision label we need for this demo is 'ground',
	// indicating something we shouldn't be able to fall or walk through
	Ground collision.Label = 1
)

func searchMinIndex(intSlice []int) int {
	min := intSlice[0]
	minIndex := 0
	for i, v := range intSlice {
		if v < min {
			min = intSlice[i]
			minIndex = i
		}
	}
	return minIndex
}

func selectionSort(intSlice []int) []int {
	sortedSlice := []int{}

	for _ = range intSlice {
		minIndex := searchMinIndex(intSlice)
		valor := intSlice[minIndex]
		intSlice = append(intSlice[:minIndex], intSlice[minIndex+1:]...)
		sortedSlice = append(sortedSlice, valor)
	}
	return sortedSlice
}

func main() {
	// slice := []int{5, 422, 3, 25, 8, 1, 11, 13}
	// fmt.Println(searchMin(slice))
	// fmt.Println(slice[5])
	fmt.Println(selectionSort([]int{500, 13, 5, 422, 3, 25, 8, 1, 11, 13}))
	// oak.AddScene("platformer", scene.Scene{Start: func(*scene.Context) {

	// 	ground := entities.NewSolid(0, 400, 500, 20,
	// 		render.NewColorBox(500, 20, color.RGBA{0, 0, 255, 255}),
	// 		nil, 0)
	// 	ground.UpdateLabel(Ground)

	// 	render.Draw(ground.R)

	// }})
	// oak.Init("platformer")
}
