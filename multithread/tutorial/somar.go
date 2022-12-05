package tutorial

/*
Código que soma dois slices de inteiros, de forma paralela usando
GoRoutines
*/

import (
	"fmt"
	"sync"
)

// função que soma dois slices de inteiros
// num_a é o lado esquerdo da soma
// num_b é o lado direito da soma
// num_resultado é o que guarda o resultado
// inicio é o índice inicial a partir do qual começar
// fim é o índice que deve parar
func somar(num_resultado []int,num_a []int,num_b []int,inicio int, fim int) {
	for i := inicio; i< fim; i++ {
		soma := num_a[i] + num_b[i]
		num_resultado[i] = soma
	}
	wg.Done() // Diz que terminou de contar
}

// O WaitGroup é usado para esperar outras GoRoutines
var wg2 sync.WaitGroup

func Main2() {
	// Cria uma lista com 4000 números
	var tamanho = 4000
	var a []int = make([]int, tamanho)
	var b []int  = make([]int, tamanho)
	var resultado []int  = make([]int, tamanho)

	// Inicia com valores quaisquer, no caso a soma deles é igual a tamanho
	// então se o resultado ficar tudo 4000, 4000, 4000 é porque deu certo
	for i := 0; i< tamanho; i++ {
		a[i] = i
		b[i] = tamanho-i
	}

	wg2.Add(4) // Diz que vai Adicionar 4 threads no wait group

	// chama a função de somar como uma GoRoutine
	// cada função irá somar uma parte diferente da lista
	go somar(resultado,a,b,0,1000)
	go somar(resultado,a,b,1000,2000)
	go somar(resultado,a,b,2000,3000)
	go somar(resultado,a,b,3000,4000)

	wg2.Wait() // Espera todos informarem que terminaram com wg.Done()
	fmt.Println("Soma encerrada")

	fmt.Println("10 primeiros valores:")
	for i := 0; i< 10; i++ {
		fmt.Print(resultado[i],", ")
	}
	fmt.Println()
}