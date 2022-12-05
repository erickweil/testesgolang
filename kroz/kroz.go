package main

import (
	"example/inputhelper"
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

type Sala struct {
	nome      string 
	descricao string
	conexoes  [12]string
}

var COMANDOS []string
var SALAS [20]Sala
var SALASCONN []int

// Largura do terminal do escreva
var LARG int = 60

// Variáveis
var posicao = 0 // Sala atual

func escreva(txt string) {
	var procurandoFim bool = false
	for i, c := range txt {
		fmt.Printf("%c", c)

		if (i+1)%LARG == 0 {
			procurandoFim = true
		}

		if procurandoFim && (unicode.IsSpace(c) || unicode.IsPunct(c)) {
			procurandoFim = false
			fmt.Println()
		}
		//time.Sleep(100)
	}
}

var criarSalaI int = 0

func criarSala(nome string, descricao string, conexoes string) {
	var s Sala
	s.nome = nome
	s.descricao = descricao
	var conn_map [12]string
	var conn_str []string = strings.Split(conexoes, "|")
	for _, str := range conn_str {
		var parts []string = strings.Split(str, " ")
		var index = encontrarcomando(parts[0])

		if index >= 0 && index < 12 {
			conn_map[index] = parts[1]
		}
	}

	s.conexoes = conn_map
	SALAS[criarSalaI] = s
	criarSalaI++
}

func inicializar() {
	inputhelper.InputInit()
	posicao = 1

	COMANDOS = []string{
		// AÇÕES DE MOVIMENTAÇÃO
		"N", "NORTE", //0
		"S", "SUL", //1
		"L", "LESTE", //2
		"O", "OESTE", //3
		"NO", "NOROESTE", //4
		"SE", "SUDESTE", //5
		"NE", "NORDESTE", //6
		"SO", "SUDOESTE", //7
		"C", "SUBIR,CIMA,ACIMA,SOBE,ERGUER,LEVANTAR", //8
		"D", "DESCER,ABAIXO,ABAIXAR", //9
		"E", "ENTRAR,DENTRO", //10
		"F", "SAIR,FORA", //11
	}

	criarSala("NULL", "", "")
	criarSala("CASAO", "Você está em uma área aberta oeste de uma casa branca, A porta da frente está fechada.", "")
	criarSala("CASAN", "Você está voltado para o lado norte de uma casa branca. Não há portas aqui, e todas as janelas estão fechadas. Ao norte um caminho estreito vai por entre as arvores", "")
	criarSala("CASAL", "Você está nos fundos de uma casa branca. Um caminho leva à floresta ao leste.", "")
	criarSala("CASAS", "Você está voltado ao lado sul de uma casa branca. Não há portas aqui, e todas as janelas estão fechadas.", "")
	criarSala("CLAR", "Você está em uma pequena clareira em um caminho na floresta bem feito que se estende leste-oeste", "")
	criarSala("F1", "Essa é uma floresta bem escura, com árvores altas em todas as direções.", "")
	criarSala("F2", "Essa é a floresta, com árvores em todas as direções. Ao leste, parece ter luz do sol", "")
	criarSala("CAMINHO", "Esse é um caminho que se passa por uma floresta escura.O caminho leva norte-sul. Uma árvore beem grande com alguns galhos baixos está próxima do caminho", "")

	/*
		N  S  L  O  NO SE NE SO C  D  E  F
	*/
	SALASCONN = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // Null
		2, 4, 0, 7, 0, 4, 2, 0, 0, 0, 0, 0, // 1 - Oeste casa
		8, 0, 3, 1, 0, 3, 0, 1, 0, 0, 0, 0, // 2 - Norte casa
		2, 4, 5, 0, 2, 0, 0, 4, 0, 0, 0, 0, // 3 - Leste casa
		0, 6, 3, 1, 1, 0, 3, 0, 0, 0, 0, 0, // 4 - Sul casa
		0, 6, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, // 5 - Clareira
		5, 0, 0, 7, 4, 0, 0, 0, 0, 0, 0, 0, // 6 - Floresta 1
		0, 6, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 7 - Floresta 2
		0, 2, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, // 8 - Caminho
	}
}

func encontrarcomando(cmd string) int {
	cmd = strings.ToUpper(cmd)
	for i := 0; i < len(COMANDOS); i += 2 {
		comando := COMANDOS[i]
		if comando == cmd {
			return i
		}

		aliases := strings.Split(COMANDOS[i+1], ",")
		for _, alias := range aliases {
			if alias == cmd {
				return i
			}
		}
	}
	return -1
}

func jogoinput() {
	fmt.Print(">")
	//fmt.Scan(&texto)
	inputhelper.InitParser(inputhelper.Readline())

	//partes := strings.Split(texto, " ")
	var arg0 string = inputhelper.NextArgument()
	cmd := encontrarcomando(arg0)

	if cmd >= 0 && cmd < len(COMANDOS) {
		fmt.Println("COMANDO:", COMANDOS[cmd])
	}
	// Movimentação
	if cmd >= 0 && cmd <= 11 {
		sala := SALASCONN[posicao*12+cmd]

		if sala >= 1 && sala < len(SALAS) {
			posicao = sala
		} else {
			escreva("Não é possível ir nesta direção \n")
		}
	} else {
		var rdn int = rand.Intn(15)

		switch rdn {
		case 0:
			escreva("Comando desconhecido\n")
		case 1:
			escreva("Precisa de Ajuda? escreva AJUDA\n")
		case 2:
			escreva("Me desculpe?\n")
		case 3:
			escreva("Não entendi, experimente:Ir norte\n")
		case 4:
			escreva("Você precisa dar uma ordem clara\n")
		case 5:
			escreva("Não é assim que você vai ganhar esse jogo\n")
		case 6:
			escreva("Tem certeza que não precisa de ajuda?\n")
		case 7:
			escreva("Você tem orgulho disso?\n")
		case 8:
			escreva("... ?\n")
		case 9:
			escreva("VOCÊ MORREU!\n\n... brincadeirinha, dê um ordem correta da próxima vez.\n")
		case 10:
			escreva("O que você quer fazer?\n")
		case 11:
			escreva("Tente fazer sentido\n")
		case 12:
			escreva("Você está realmente achando que vou entender isso?\n")
		case 13:
			escreva("sflhs fsfh sd fjsdhfsdjfsdfsd sdfsdf\nFoi isso que entendi\n")
		case 14:
			escreva("Não\n")
		default:
			escreva("Nada acontece\n")
		}
	}
}

func descrever() {
	var nome string = SALAS[posicao].nome
	escreva(nome)
	escreva("\n")

	var desc string = SALAS[posicao].descricao
	escreva(desc)
	escreva("\n")
}

func main() {
	inicializar()

	for {
		descrever()
		jogoinput()
	}
}
