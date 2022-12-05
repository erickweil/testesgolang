package tutorial

/*
Código que encontra o maior número em uma lista

Já que é necessário analisar o retorno da GoRoutine este
exemplo utiliza channels para receber os valores, dispensando
o uso de WaitGroups
*/

import (
	"fmt"
)

// função que encontra o maior número numa Slice
// A ideia é que retorne o maior número no channel
func maior(nums []int, channel chan int) {
	// escolhe o chute como sendo o primeiro numero da lista
	var chute int = nums[0]

	//PARA CADA NÚMERO DA LISTA
	for i := 0; i< len(nums); i++ {
		if nums[i] > chute {
			// SE ACHOU UM NÚMERO MAIOR
			// SUBSTITUI O CHUTE
			chute = nums[i]
		}
	}
	// ao final deste loop o valor de chute será o maior
	// número do slice

	// Em vez de retornar,
	// o resultado da GoRoutine é enviado no channel
	channel <- chute
}

func Main3() {
	// Cria uma lista com 4000 números
	var tamanho = 4000
	var nums []int = make([]int, tamanho)

	// Inicia com valores quaisquer, nesse caso o maior número será 3999
	for i := 0; i< tamanho; i++ {
		nums[i] = i
	}


	// cria um channel, isto será usado para comunicação
	// entre as GoRoutines
	channel := make(chan int)

	// Encontrar o maior usando 4 GoRoutines
	// Fazendo slicing da lista para passar pedaços diferentes dos números
	// a serem analisado por cada
	go maior(nums[0:1000],channel)
	go maior(nums[1000:2000],channel)
	go maior(nums[2000:3000],channel)
	go maior(nums[3000:4000],channel)

	// Recebe os 4 resultados
	// No caso irá esperar cada GoRoutine terminar e retornar seus resultados
	// Não necessariamente estará em ordem
	maior1000 := <-channel
	maior2000 := <-channel
	maior3000 := <-channel
	maior4000 := <-channel

	// Calcula o maior entre os 4 sub-resultados
	maiorTodos := maior1000
	if maior2000 > maiorTodos { maiorTodos = maior2000 }
	if maior3000 > maiorTodos { maiorTodos = maior3000 }
	if maior4000 > maiorTodos { maiorTodos = maior4000 }

	fmt.Println("Calculo encerrado")

	fmt.Println("Maior de todos:")
	fmt.Println(maiorTodos)
}