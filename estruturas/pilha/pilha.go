package pilha

import "fmt"

type Pilha struct {
	lista []int
	topo  int
}

func (pilha *Pilha) push(valor int) {
	pilha.lista[pilha.topo] = valor
	pilha.topo++
}

func (pilha *Pilha) pop() int {
	pilha.topo--
	return pilha.lista[pilha.topo]
}

func (pilha *Pilha) isEmpty() bool {
	return pilha.topo == 0
}

func Main() {
	var pilha = Pilha{lista: make([]int, 100)}

	pilha.push(10)
	pilha.push(20)
	pilha.push(30)

	fmt.Println(pilha.pop())
	fmt.Println(pilha.pop())
	fmt.Println(pilha.pop())
}