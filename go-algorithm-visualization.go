package main

import (
	"fmt"
	"image/color"
	"math"
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

func searchMinIndexCB(CBSlice []*render.Sprite) int {
	min := CBSlice[0].GetRGBA().Rect.Min.Y
	minIndex := 0
	for i, v := range CBSlice {
		if math.Abs(float64(v.GetRGBA().Rect.Min.Y)) < math.Abs(float64(min)) {
			min = CBSlice[i].GetRGBA().Rect.Min.Y
			minIndex = i
		}
	}
	return minIndex
}

func selectionSortCB(CBSlice []*render.Sprite) []*render.Sprite {
	sortedSlice := []*render.Sprite{}

	for _, _ = range CBSlice {
		minIndex := searchMinIndexCB(CBSlice)
		colorBox := CBSlice[minIndex]
		// fmt.Printf("X: %+v\n", colorBox.Vector.X())
		// fmt.Printf("Y: %+v\n", colorBox.Vector.Y())
		// fmt.Printf("X RGBA: %+v\n", colorBox.GetRGBA().Rect.Min.X)
		// fmt.Printf("Y RGBA: %+v\n", colorBox.GetRGBA().Rect.Min.Y)
		CBSlice = append(CBSlice[:minIndex], CBSlice[minIndex+1:]...)
		sortedSlice = append(sortedSlice, colorBox)
	}
	fmt.Println("Tamanho", len(sortedSlice))
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
			// for i := 0; i < len(colorBoxes); i++ {
			// 	fmt.Printf("X: %+v\n", colorBoxes[i].Vector.X())
			// 	fmt.Printf("Y: %+v\n", colorBoxes[i].Vector.Y())
			// 	fmt.Printf("X RGBA: %+v\n", colorBoxes[i].GetRGBA().Rect.Min.X)
			// 	fmt.Printf("Y RGBA: %+v\n", colorBoxes[i].GetRGBA().Rect.Min.Y)
			// }
			cbs := selectionSortCB(colorBoxes)
			for i := 0; i < len(cbs); i++ {
				fmt.Printf("X: %+v\n", cbs[i].Vector.X())
				fmt.Printf("Y: %+v\n", cbs[i].Vector.Y())
				fmt.Printf("X RGBA: %+v\n", cbs[i].GetRGBA().Rect.Min.X)
				fmt.Printf("Y RGBA: %+v\n", cbs[i].GetRGBA().Rect.Min.Y)
			}
			for i := 0; i < len(cbs); i++ {
				drawColorBoxes(ctx, cbs[i])
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
