package listaencadeada

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type ListaEncadeada struct {
	inicio *No
	fim    *No
}

func (lista *ListaEncadeada) Add(v int) {
	var novoNo = No{valor: v, proximo: nil}
	if lista.fim == nil {
		lista.inicio = &novoNo
		lista.fim = &novoNo
	} else {
		lista.fim.proximo = &novoNo
		lista.fim = &novoNo
	}
}

func (lista *ListaEncadeada) getno(indice int) *No {
	return nil
}


func (lista *ListaEncadeada) Get(indice int) int {
	return lista.getno(indice).valor
}

func Main() {
	var lista = ListaEncadeada{}
	lista.Add(10)
	lista.Add(20)

	fmt.Println(lista)
	fmt.Println("Inicio:",lista.inicio)
	fmt.Println("Fim:",lista.fim)
}