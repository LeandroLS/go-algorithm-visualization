package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/oakmound/oak/v3"
	"github.com/oakmound/oak/v3/render"
	"github.com/oakmound/oak/v3/scene"
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

const qtdColorBox = 100

func makeColorBoxes(qtd int) []*render.Sprite {
	colorBoxes := []*render.Sprite{}
	for i := 0; i < qtd; i++ {
		randInt := rand.Intn(100-10) + 10
		cb := render.NewColorBoxM(1, -randInt, color.White)
		cb.SetPos(float64(i), 100)
		colorBoxes = append(colorBoxes, cb)
	}
	return colorBoxes
}

func drawColorBoxes(ctx *scene.Context, colorBox *render.Sprite) {
	ctx.DrawStack.Draw(colorBox, 0)
}

func main() {
	c1 := oak.NewWindow()
	c1.DrawStack = render.NewDrawStack(render.NewDynamicHeap())
	c1.FirstSceneInput = color.RGBA{255, 0, 0, 255}
	c1.AddScene("scene1", scene.Scene{
		Start: func(ctx *scene.Context) {
			fmt.Println("Start scene 1")
			colorBoxes := makeColorBoxes(qtdColorBox)
			for i := 0; i < len(colorBoxes); i++ {
				drawColorBoxes(ctx, colorBoxes[i])
			}
			dFPS := render.NewDrawFPS(0.1, nil, 600, 10)
			ctx.DrawStack.Draw(dFPS, 1)
		},
	})
	c1.Init("scene1", func(c oak.Config) (oak.Config, error) {
		c.Debug.Level = "VERBOSE"
		c.DrawFrameRate = 1200
		c.FrameRate = 60
		c.EnableDebugConsole = true
		return c, nil
	})
}
