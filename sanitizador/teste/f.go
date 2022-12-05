package main

import "fmt"

func main() {
	var acesso int
	var usuario bool
	senha := [10]int{112233, 221133, 332211, 331122, 123456, 654321, 456123, 321654, 415263, 142536}
	fmt.Print("Informe sua senha:")
	fmt.Scan(&acesso)

	for _, senhas := range senha {
		if senhas == acesso {
			usuario = true
			break
		} else {
			usuario = false
		}
	}
	if usuario == true {
		fmt.Println("Acesso Concedido!")
	} else {
		fmt.Println("Senha inválida!")
	}
}