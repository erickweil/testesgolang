package aleatorio

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
Código que gera uma sequência aleatória de texto
que então é escrita em um arquivo.
*/

/*
Fnção que salva um texto em um arquivo
Fonte: https://golangbot.com/write-files/
*/
func salvarArquivo(texto string, caminho string) {
	f,err := os.Create(caminho)
	if err != nil {
		panic("NÃO CONSEGUIU CRIAR ARQUIVO")
	}
	defer f.Close()

	n,err := f.WriteString(texto)
	if err != nil {
		panic("NÃO CONSEGUIU ESCREVER ARQUIVO")
	}

	fmt.Println("Escritos",n,"caracteres com sucesso!")
}

// Gerar uma sequência aleatória de texto
func gerarTexto(tamanho int, maxlinha int, alfabeto []rune) []rune {
	// inicia o random com uma seed
	// A seed é escolhida como o tempo atual
	// assim a cada execução a seed será outra
	rand.Seed(time.Now().Unix()) 

	// Cria o texto como um slice de rune
	// Desta forma pode alterar o valor dos índices que quiser
	// No GoLang um rune é como um número mas que representa uma letra
	var texto []rune = make([]rune, tamanho)
	// calcula o tamanho do alfabeto aqui
	var alfabeto_len int = len(alfabeto)
	for i := 0; i < tamanho; i++ {
		var letra rune
		// pular linha quando chegar no tamanho máximo da linha
		if i % maxlinha == 0 && i != 0{
			letra = '\n'
		} else {
			// calcula um índice aleatório para escolher
			// uma letra dentro do alfabeto
			var indice int = rand.Intn(alfabeto_len)
			letra = alfabeto[indice]
		}

		// seta o valor da letra na posição 'i'
		texto[i] = letra
	}

	return texto
}

func Main() {
	main()
}

func main() {
	// Letras que poderão ser usadas na geração
	var alfabeto []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 ")
	
	start := time.Now() // Medindo tempo

	var texto = gerarTexto(10000000,80,alfabeto)

	elapsed := time.Since(start)
	fmt.Printf("Demorou %s para gerar o texto aleatório\n", elapsed)

	salvarArquivo(string(texto),"./aleatorio/aleatorio.txt")
}