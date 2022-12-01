package main

import "math"

type Vector2 struct {
	x float64
	y float64
}

func (a Vector2) somar(b Vector2) Vector2 {
	return Vector2{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (a Vector2) subtrair(b Vector2) Vector2 {
	return Vector2{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a Vector2) magnitude() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}

func (a Vector2) distanciaAte(b Vector2) float64 {
	var C Vector2 = b.subtrair(a)
	return C.magnitude()
}