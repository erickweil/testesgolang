package main

import (
	"fmt"
	"strings"
)

func maior(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func repete(txt string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(txt)
	}
}

func validaEmail(email string) (bool, string) {
	if !strings.Contains(email,"@") { 
		return false,"Não contém arroba!" 
	}
	if !strings.Contains(email,".") { 
		return false,"Não possui ponto" 
	}
	if len(email) < 3 { 
		return false,"Menor que 3 caracteres"
	}

	return true,""
}

func main() {
	var numeros = []int{1,2,3,4,5,6,7,8,9}

	for _, num := range numeros {
		fmt.Println(num)
	}
}