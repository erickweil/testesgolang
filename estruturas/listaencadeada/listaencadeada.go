package listaencadeada

import "fmt"

type No struct {
	valor int
	proximo *No
}

type ListaEncadeada struct {
	inicio *No
	fim *No
}

func (lista ListaEncadeada) Add(valor int) {
	fmt.Println("Add",valor)
}

func Main() {
	var lista = ListaEncadeada{}
	lista.Add(10)
}