package inputhelper

import (
	"strings"
)

func FindExact(arg string,lista []string) int {
	for i := 0; i < len(lista); i++ {
		if lista[i] == arg {
			return i
		}
	}
	return -1	
}

func FindIgnoreCase(arg string,lista []string) int {
	ARG := strings.ToUpper(arg)
	for i := 0; i < len(lista); i++ {
		A := strings.ToUpper(lista[i])
		if A == ARG {
			return i
		}
	}
	return -1	
}

// Retorna o quão parecidas são as duas strings
func CalcLikeness(a string, b string) float32 {
	// 1.0 quando iguais
	if a == b { return 1.0 }
	
	A := strings.ToUpper(a)
	B := strings.ToUpper(b)
	if A == B { return 0.999 }

	// comparar letras
	A_runes := []rune(A)
	B_runes := []rune(B)
	Asize := len(A_runes)
	Bsize := len(B_runes)
	off := make([]int, Asize)

	// calcular offset das letras
	for i := 0; i < Asize; i++{
		off[i] = -1
		for k := 0; k < Bsize; k++{
			if A_runes[i] == B_runes[k] {
				off[i] = k
			}
		}
	}

	// procurar a maior quantidade de offsets consecutivos iguais diferentes de -1
	//var count int = 0
	//for i := 0; i < Asize; i++{

	//}
	
	return 0
}