package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var posTri = Vector2{x:400.0,y:300.0} 
var velTri = Vector2{x:0,y:0}
var aceTri = Vector2{x:0,y:-0.01}

func draw(imd *imdraw.IMDraw) {

	velTri = velTri.somar(aceTri)
	posTri = posTri.somar(velTri)


	imd.Color = pixel.RGB(1, 0, 0)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(50+posTri.x, 25+posTri.y))

	imd.Color = pixel.RGB(0, 1, 0)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(200+posTri.x, 25+posTri.y))
	
	imd.Color = pixel.RGB(0, 0, 1)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(125+posTri.x, 175+posTri.y))

	
	imd.Color = pixel.RGB(1, 0, 0)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(50+posTri.x, 25+posTri.y))
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