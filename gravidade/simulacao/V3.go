package simulacao

import (
	"math"
	"math/rand"
)

type V3 struct {
	X float64
	Y float64
	Z float64
}

func (a V3) Add(b V3) V3 {
	return V3{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func (a V3) Sub(b V3) V3 {
	return V3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (a V3) Mul(b float64) V3 {
	return V3{X: a.X * b, Y: a.Y * b, Z: a.Z * b}
}

func (a V3) Div(b float64) V3 {
	return V3{X: a.X / b, Y: a.Y / b, Z: a.Z / b}
}

func (a V3) SqrMagnitude() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a V3) Magnitude() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func V3_rand(min float64, max float64) V3 {
	scale := max - min
	return V3{
		X: ((rand.Float64()-0.5) * scale) + min,
		Y: ((rand.Float64()-0.5) * scale) + min,
		Z: ((rand.Float64()-0.5) * scale) + min,
	}
}


