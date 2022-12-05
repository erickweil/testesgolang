package simulacao

type Objeto struct {
	P   V3
	V   V3
	M   float64
	R   float64
	Cor [3]uint8
}

var G = 1.0

func (obj *Objeto) atrair(outro *Objeto, dt float64) {
	var dir = obj.P.Sub(outro.P)
	var dist = dir.Magnitude()

	if dist < 0.000001 {
		return
	}

	dir = dir.Div(dist) // normalizar

	if dist < obj.R+outro.R {
		dist = obj.R + outro.R
	}

	var Fmag = G * (obj.M * outro.M) / (dist * dist)
	var F = dir.Mul(-Fmag)

	// F = m * a
	// a = F / m
	var A = F.Div(obj.M)

	obj.V = obj.V.Add(A.Mul(dt))
}

func (obj *Objeto) etapaFisica(dt float64) {
	obj.P = obj.P.Add(obj.V.Mul(dt))

}