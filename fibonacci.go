package main

import (
	"math/rand"
	"time"
	"image/color"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Fibonacci",
		Bounds: pixel.R(0, 0, 720, 420),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// create draw pen
	imd := imdraw.New(nil)


	
	colourR := rand.Intn(256)
	colourG := rand.Intn(256)
	colourB := rand.Intn(256)

	x := 0.0
	y := 340.0

	for !win.Closed() {
		win.Clear(colornames.Black)

		imd.Color = color.RGBA{	uint8(adjustColourValue(colourR)),
								uint8(adjustColourValue(colourG)),
								uint8(adjustColourValue(colourB)), 0xff }
		imd.Push(pixel.V(x, y))
		imd.Ellipse(pixel.V(3, 3), 0)

		imd.Draw(win)
		win.Update()
		x = x + 1
		if x == 700 {
			x = 0
			y -= 10
		}
		// fmt.Printf("%s\n", x)
	}
}

func adjustColourValue(currentValue int) int {
	fmt.Printf("%s\n", currentValue)
	if currentValue > 254  {
		currentValue = 0
	} else {
		currentValue += 1
	}
	fmt.Printf("%s\n", currentValue)
	return currentValue
}

// func adjustColourValue(currentValue int) int {
// 	// 50/50 chance to change the value
// 	if rand.Float32() < 0.5 {
// 		// if range colour is maxed or minned, then force it the other way
// 		if currentValue == 256 {
// 			currentValue -= 10
// 		} else if currentValue == 0{
// 			currentValue += 10
// 		} else {
// 			// or do another coin toss
// 			if rand.Float32() < 0.5 {
// 				currentValue += 10
// 			} else {
// 				currentValue -= 10
// 			}
// 		}
// 	}
// 	fmt.Printf("%s\n", currentValue)
// 	return currentValue
// }

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}