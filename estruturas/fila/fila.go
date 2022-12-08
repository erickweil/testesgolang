package fila

import (
	"fmt"
)

type Fila struct {
	lista []int
	inicio  int
	fim  int
}

func (f *Fila) enqueue(valor int) {
	f.lista[f.fim] = valor
	f.fim = (f.fim + 1) % len(f.lista)
}

func (f *Fila) dequeue() int {
	var ret = f.lista[f.inicio]
	f.inicio = (f.inicio + 1) % len(f.lista)
	return ret
}

func (f *Fila) isEmpty() bool {
	return f.inicio == f.fim
}

func Main() {
	var fila = Fila{lista: make([]int, 100)}

	fila.enqueue(10)
	fila.enqueue(20)
	fila.enqueue(30)

	fmt.Println(fila.dequeue())
	fmt.Println(fila.dequeue())
	fmt.Println(fila.dequeue())
}