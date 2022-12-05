package simulacao

import (
	"math"
)

var Objs []Objeto

func IniciarSimulacao() {
	Objs = []Objeto{
		{
			P:   V3{X: 0, Y: 0, Z: 0},
			V:   V3{X: 0, Y: 0, Z: 0},
			M:   10000.0,
			Cor: [3]uint8{255, 0, 0},
		},
		
		{
			P:   V3{X: 0, Y: 300, Z: 0},
			V:   V3{X: -6.15, Y: 0.1, Z: 0},
			M:   1000.0,
			Cor: [3]uint8{0, 120, 255},
		},
		{
			P:   V3{X: 0, Y: 350, Z: 0},
			V:   V3{X: -1.50, Y: 0.1, Z: 0},
			M:   50.0,
			Cor: [3]uint8{100, 0, 0},
		},
	}

	/*Objs = make([]Objeto, 100)
	p_max := 100.0
	v_max := 3.0
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < len(Objs); i++ {
		Objs[i] = Objeto{
			P:   V3_rand(-p_max,p_max),
			V:   V3_rand(-v_max,v_max),
			M:   rand.Float64() * rand.Float64() * 1000.0,
			Cor: [3]uint8{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))},
		}
	}*/

	for i := 0; i < len(Objs); i++ {
		// Calcular Raio baseado na massa
		// https://www.omnicalculator.com/physics/sphere-density
		var obj = &Objs[i]

		// (4/3) × π × r3 = mass/density
		// r3 = (mass/density) / ((4/3) × π)
		// r = Sqrt3((mass/density) / ((4/3) × π))

		var massdensity = obj.M / 0.05520
		var pi43 = (4.0/3.0) * math.Pi
		var r3 = massdensity / pi43

		obj.R = math.Cbrt(r3)

		if obj.R > 100 {
			obj.R = 100
		}
	}
}

var Centro V3

func SimularGravidade(etapas int, tempo float64) {
	for i:=0;i<etapas;i++ {
		EtapaGravidade(tempo/float64(etapas))
	}
}
func EtapaGravidade(dt float64) {
	for x := 0; x < len(Objs); x++ {
		for y := 0; y < len(Objs); y++ {
			if x == y {
				continue
			}

			var ObjA = &Objs[x]
			var ObjB = &Objs[y]
			ObjA.atrair(ObjB,dt)
		}
	}
	Centro = V3{X: 0.0, Y: 0.0, Z: 0.0}
	var totalMassa = 0.0
	for i := 0; i < len(Objs); i++ {
		Objs[i].etapaFisica(dt)
		Centro = Centro.Add(Objs[i].P.Mul(Objs[i].M))
		totalMassa += Objs[i].M
	}
	Centro = Centro.Div(totalMassa)
}