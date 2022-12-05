package raymarch

import "math"

/*
Cálculos do raymarch
NÃO PRECISA MEXER AQUI

https://jamie-wong.com/2016/07/15/ray-marching-signed-distance-functions/
https://iquilezles.org/
	https://iquilezles.org/articles/distfunctions/

*/
const MAXIMORAYMARCH = 1024
const MAXDIST = 20.0
const EPSILON = 0.0001
const NORMEPS = (MAXDIST/float64(MAXIMORAYMARCH)) / 100.0

func esfera(pos V3,raio float64) float64 {
	return pos.Magnitude() - raio
}

func plano(pos V3, norm V3, h float64) float64 {
	return pos.Dot(norm) + h
}

func torus(pos V3, t V3) float64 {
	q := V3{
		X:V3{pos.X,pos.Y,0.0}.Magnitude()-t.X,
		Y:pos.Z,
		Z:0.0,
	}

	return q.Magnitude()-t.Y
}

func distancia(p V3) (dist float64,refletir bool,cor V3) {
	chao := plano(p,V3{0,-1.0,0},1.0)
	prato := math.Max(
		-esfera(p.Sub(V3{0.8,-0.2,3.0}),1.35),
		esfera(p.Sub(V3{0.8,-0.1,3.0}),1.3),
	)
	ret := math.Min(chao,prato)
	
	laranja1 := esfera(p.Sub(V3{0.5,0.7,3.2}),0.4)
	laranja2 := esfera(p.Sub(V3{1.3,0.7,3.0}),0.35)
	laranja3 := esfera(p.Sub(V3{0.6,0.7,3.4}),0.45)
	laranja4 := esfera(p.Sub(V3{1.0,0.7,2.6}),0.3)
	laranja5 := esfera(p.Sub(V3{1.0,0.25,3.05}),0.45)
	laranjas := math.Min(laranja1,laranja2)
	laranjas = math.Min(laranjas,laranja3)
	laranjas = math.Min(laranjas,laranja4)
	laranjas = math.Min(laranjas,laranja5)
	ret = math.Min(ret,laranjas)

	corte := -plano(p.Sub(V3{-1.8,0.65,3.2}),V3{-0.8,1.6,0.3}.Normalized(),0.1)
	laranjaantes := esfera(p.Sub(V3{-1.8,0.65,3.2}),0.4)
	laranjacortada := math.Max(
		corte,
		laranjaantes,
	)

	ret = math.Min(ret,laranjacortada)

	laranjarefl := esfera(p.Sub(V3{-0.81,0.65,3.3}),0.35)
	ret = math.Min(ret,laranjarefl)

	if ret == laranjarefl {
		refletir = true
	} else {
		refletir = false
	}

	if ret == chao {
		cor = V3{0.6,0.45,0.47}
	} else if ret == prato {
		cor = V3{1.0,1.0,1.0}
	} else if ret == laranjas {
		cor = V3{0.88,0.59,0.15}
	} else if ret == corte {
		if math.Abs(laranjaantes) < 0.05 || math.Abs(laranjaantes) > 0.27 {
			cor = V3{1.0,1.0,1.0}
		} else {
			cor = V3{0.88,0.59,0.15} 
		}
	} else if ret == laranjaantes {
		cor = V3{0.88,0.59,0.15} 
	} else {
		cor = V3{1.0,0.0,1.0}
	}

	dist = ret
	return
}

func normal(p V3) V3 {
	x1, _, _ := distancia(p.Add(V3{NORMEPS,0.0,0.0}))
	x2, _, _ := distancia(p.Sub(V3{NORMEPS,0.0,0.0}))
	y1, _, _ := distancia(p.Add(V3{0.0,NORMEPS,0.0}))
	y2, _, _ := distancia(p.Sub(V3{0.0,NORMEPS,0.0}))
	z1, _, _ := distancia(p.Add(V3{0.0,0.0,NORMEPS}))
	z2, _, _ := distancia(p.Sub(V3{0.0,0.0,NORMEPS}))
	var gradiente = V3{
		x1 - x2,
		y1 - y2,
		z1 - z2,	
	}
	return gradiente.Normalized()
}

func calcLambert(p V3, normal V3, light V3) float64 {
	lightDirection := light.Sub(p)
	lightDist := lightDirection.Magnitude()
	lightDirection = lightDirection.Div(lightDist)

	d,_ := fazerRayMarch(p.Add(normal.Mul(EPSILON*10.0)),lightDirection,0.01)
	if d < lightDist && d < MAXDIST {
		return 0.15
	} else {
		return math.Max(0.20,lightDirection.Dot(normal))
	}
}

var refletiu = false
func fazerRayMarch(pixel V3, raioDir V3, profundidade float64) (float64,V3) {
	//olho := V3{X:0.0,Y:0.0,Z:0.0}
	var cor V3
	for i := 0; i < MAXIMORAYMARCH; i++ {
		d,refletir,_cor := distancia(pixel.Add(raioDir.Mul(profundidade)))
		cor = _cor

		if d < EPSILON {

			if refletir && !refletiu {
				final_p := pixel.Add(raioDir.Mul(profundidade))
				norm := normal(final_p)
				refl := raioDir.Sub(norm.Mul(raioDir.Dot(norm)).Mul(2.0))
				refletiu = true
				return fazerRayMarch(final_p, refl, 0.01)
			} else {
				return profundidade,cor
			}
		}

		profundidade += d

		if profundidade >= MAXDIST {
			return profundidade,cor
		}
	}

	return profundidade,cor
}

func renderizarPixel(x float64, y float64) V3 {
	p := V3{
		X:x,
		Y:y,
		Z:1.0, // NEAR PLANE, ~45º
	}
	profundidade := p.Magnitude()
	raioDir := p.Div(profundidade)
	refletiu = false
	d,cor := fazerRayMarch(V3{0,0,0},raioDir,profundidade)

	if d < MAXDIST {
		final_p := raioDir.Mul(d)
		norm := normal(final_p)
		light := calcLambert(final_p,norm,V3{10.0,-10.0,-10})
		//obj_cor := cor(final_p)
		return cor.Mul(light)
	} else {
		return V3{
			X:0.65,
			Y:0.88,
			Z:0.92,
		}
	}
}

