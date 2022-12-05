package tutorial

/*
Código que conta até 100 em duas GoRoutines
*/

import (
	"fmt"
	"sync"
	"time"
)

// função que conta até o valor especificado
// contar("a",3):
//	a: 0
//	a: 1
//	a: 2
func contar(nome string,quantos int) {
	for i := 0; i< quantos; i++ {
		fmt.Print(nome,":",i,"\n")
		time.Sleep(time.Second) // dorme 1 segundo
	}
	wg.Done() // Diz que terminou de contar
}

// O WaitGroup é usado para esperar outras GoRoutines
var wg sync.WaitGroup

func Main1() {

	wg.Add(2) // Diz que vai Adicionar 2 threads no wait group

	// chama a função contar como uma GoRoutine
	// A palavrinha 'go' antes que faz isso
	go contar("a",10)
	go contar("b",10)

	wg.Wait() // Espera todos informarem que terminaram com wg.Done()
	fmt.Println("Contagem encerrada")
}