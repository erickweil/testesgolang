package senha

/*
Código que pesquisa todas as combinações de dígitos
para descobrir qual senha produz o hash md5 informado.

É basicamente um brute-force para reverter um hash md5

A pesquisa é limitada a senhas de 8 dígitos numéricos
de forma a simplificar o código e permitir que o mesmo
termine antes do sol engolir a terra e o universo inteiro acabar.

Foi escolhido o hash md5 pelo mesmo ser relativamente rápido
de se calcular, demostrando assim porque algoritmos de hash lentos
são mais seguros.

Para simplificar ainda mais não há 'salt' no hash gerado.
*/

import (
	"crypto/md5"   // Pacote que permite calcular hashes
	"encoding/hex" // Para converter o hash para hexadecimal
	"fmt"
	"time"
)

/*
Função que calcula o hash de um texto e retorna um texto em hexadecimal
fonte: https://stackoverflow.com/questions/2377881/how-to-get-a-md5-hash-from-a-string-in-golang
*/
func HashMD5(txt string) string {
	var hash [16]byte = md5.Sum([]byte(txt)) // Calcula o hash do texto, o que produz um array de 16 bytes
	var hash_hex string = hex.EncodeToString(hash[:]) // Converte o array de 16 bytes em um texto hexadecimal

	// retorna o hash em hexadecimal
	return hash_hex
 }

/*
Itera por todas as combinações de senhas de 8 dígitos e
compara o hash de cada uma com o valor de 'senha_hash'
quando encontrar a senha que gera o mesmo valor de hash, retorna
esta senha.

Se não encontrar retorna um texto vazio
*/
func Bruteforce(senha_hash string) string {
	// 1. Atravessar todas as possibilidades de números
	// desde 00000000 até 99999999
	for i := 0; i <= 99999999; i++ {
		// converte o número do contador para uma senha de 8 dígitos
		// ex. o número 123 vira o texto '00000123', 1 vira '00000001' e etc...
		// fonte: https://stackoverflow.com/questions/25637440/how-to-pad-a-number-with-zeros-when-printing
		var gerado string = fmt.Sprintf("%08d",i)

		// calcula o hash da senha gerada
		var gerado_hash string = HashMD5(gerado)

		// comparar se o hash é igual
		if gerado_hash == senha_hash {
			// Se é igual, ACHOU, RETORNA A SENHA!!!
			return gerado
		}
	}

	// Se chegou aqui é porque a senha não é de 6 dígitos numéricos
	// retorna um texto vazio
	return ""
}

// para poder chamar de fora do pacote.
func Main() {
	main()
}

// Executar e medir o tempo.
func main() {
	var senha_hash string
	fmt.Println("Insira o hash md5 em hexadecimal de uma senha")
	fmt.Println("Experimente inserir '25d55ad283aa400af464c76d713c07ad' ")
	fmt.Scan(&senha_hash)

	fmt.Println("Descobrindo que senha gera o hash",senha_hash)

	start := time.Now() // Medindo tempo
	
	var senha string = Bruteforce(senha_hash)

    elapsed := time.Since(start)
    fmt.Printf("Demorou %s \n", elapsed)

	if senha == "" {
		fmt.Println("O hash não é de uma senha uma entre 00000000 e 99999999")
	} else {
		fmt.Println("A senha que gerou este hash é:",senha)
	}
}
