package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Testes Canvas",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(200, 100))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(800, 100))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(500, 700))
	imd.Polygon(0)

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}