package main

import (
	"fmt"
	"regexp"
)

func main() {

testeSenha := regexp.MustCompile(`[a-z]+[A-Z]+[\d+]`)
senha := ""
fmt.Println("Digite sua senha:")
fmt.Scan(&senha)

if len(senha) < 8 {

fmt.Println("Menor que 8.")

} else if senha == "senha123" || senha == "12345678" || senha == "password" || senha == "11111111" || senha == "qwertyui" {

fmt.Println("Senha fraca, tente novamente.")

} else if true != testeSenha.MatchString(senha) {
fmt.Println("Senha errada. Senha não possui letras maiúsculas, minúsculas e números. Tente novamente.")
} else {
fmt.Println("Senha confirmada.")
}
}