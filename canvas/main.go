package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var diffX = 0.0

func draw(imd *imdraw.IMDraw) {

	diffX = diffX + 1.0

	if diffX > 1024 {
		diffX = -1024
	}

	imd.Color = pixel.RGB(1, 0, 0)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(200+diffX, 100))

	imd.Color = pixel.RGB(0, 1, 0)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(800+diffX, 100))
	
	imd.Color = pixel.RGB(0, 0, 1)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(500+diffX, 700))

	
	imd.Color = pixel.RGB(1, 0, 0)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(200+diffX, 100))
	//imd.Polygon(0)
	imd.Line(50.0)
	//imd.Circle(80.0,5.0)
}

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

	for !win.Closed() {
		imd.Clear()
		draw(imd)

		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}