package primos

// Código que calcula se um número é primou ou não usando várias threads
// Feito em aula

import (
	"fmt"
	"time"
)

// Função que não usa Threads, para comparar
func ePrimo2(n int64) bool {
	for i := int64(2); i < n; i++ {
		var resto = n % i
		if resto == 0 {
			return false
		}
	}
	return true
}

// Função que descobre se um número é primo,
// porém pesquisa em apenas em um trecho entre inicio e fim
func ePrimoTrecho(n int64, inicio int64, fim int64, result chan bool) {
	for i := inicio; i < fim; i++ {
		var resto = n % i
		if resto == 0 {
			result <- false
			return
		}
	}
	result <- true
	return
}

// Função que Descobre se um número é primo usando threads.
// A ideia é procurar divisores em intervalos, se nenhum deles
// encontrar divisores é porque é primo
func ePrimo(n int64) bool {
	channel := make(chan bool, 4)
	var nucleos = int64(12) // irá criar 12 threads, dividindo a pesquisa em 12 partes
	var t = n/nucleos
	go ePrimoTrecho(n, 2, t, channel)
	for i := int64(1); i < nucleos; i++ {
		go ePrimoTrecho(n, t*i, t*(i+1), channel)
	}
	
	// Se qualquer uma das funções encontrar que o número é divisível por outro
	// não é primo.
	var result = true
	for i := int64(0); i < nucleos; i++ {
		r := <-channel
		if !r {
			result = false
		}
	}

	return result
}

func Main() {
	main()
}

func main() {
	/*var n int64 = 1 << 32-1
	for i := n; true; i++ {
		fmt.Println(i,ePrimo(i))
	}*/

    start := time.Now()

	var n int64 =4294967311 // número primo muito grande para testar
	fmt.Println(n,ePrimo(n))
	//fmt.Println(n,ePrimo2(n))

	
    elapsed := time.Since(start)
    fmt.Printf("demorou %s", elapsed)
}