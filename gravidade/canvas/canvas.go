package canvas

import (
	"example/gravidade/simulacao"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel/imdraw"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.Beige)

		imd.Clear()

		simulacao.SimularGravidade(100,1)
		var c = simulacao.Centro
		c.X += 1024/2
		c.Y += 768/2
		c.Z -= 100
		for i := 0; i < len(simulacao.Objs); i++ {
			var obj = simulacao.Objs[i]
			imd.Color = color.RGBA{R: obj.Cor[0],G: obj.Cor[1], B: obj.Cor[2], A: 255}

			cent := c.Sub(obj.P)
			//proj := cent.Div(cent.Z)
			//projj := cent.Add(simulacao.V3{X:obj.R,Y:0,Z:0}).Div(cent.Z)
			//r := math.Abs(proj.X - projj.X)

			imd.Push(pixel.V(cent.X,cent.Y))
			imd.Ellipse(pixel.V(obj.R, obj.R), 0)
		}
		imd.Draw(win)

		win.Update()
	}
}

func Main() {
	simulacao.IniciarSimulacao()
	pixelgl.Run(run)
}