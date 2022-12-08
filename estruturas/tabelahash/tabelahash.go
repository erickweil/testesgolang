package tabelahash

import (
	"fmt"
	"hash/fnv"
)

// https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go
func StrHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type Tno struct {
	chave string
	valor string
	prox *Tno
}

type TabelaHash struct {
	tabela []*Tno
}

func checkChave(no *Tno,chave string) *Tno {
	for i := 0; no != nil; i++ {
		if no.chave == chave {
			return no
		}
		no = no.prox
	}
	return nil
}

func (t *TabelaHash) Put(chave string, valor string) {
	var no = Tno{chave:chave,valor:valor}
	var hash = StrHash(chave)
	var indice = int(hash % uint32(len(t.tabela)))

	if t.tabela[indice] == nil {
		t.tabela[indice] = &no
	} else {
		var v = checkChave(t.tabela[indice],chave)
		if v == nil {
			no.prox = t.tabela[indice]
			t.tabela[indice] = &no
		} else {
			v.valor = valor
		}
	}
	
}

func (t *TabelaHash) Get(chave string) (string,error) {

	var hash = StrHash(chave)
	var indice = int(hash % uint32(len(t.tabela)))

	if t.tabela[indice] == nil {
		return "", fmt.Errorf("Não possui esta chave")
	} else {
		var v = checkChave(t.tabela[indice],chave)
		if v == nil {
			return "", fmt.Errorf("Não possui esta chave")
		} else {
			return v.valor,nil
		}
	}
}

func Main() {
	var tabelaHash = TabelaHash{tabela: make([]*Tno, 100)}

	for i := 0;i < 1000; i++ {
		tabelaHash.Put(fmt.Sprintf("Chave %d",i),fmt.Sprintf("Valor %d",i))		
	}

	valor1, _ := tabelaHash.Get("Chave 1")
	valor2, _ := tabelaHash.Get("Chave 500")
	valor3, _ := tabelaHash.Get("Chave 333")

	fmt.Println("1:",valor1)
	fmt.Println("500:",valor2)
	fmt.Println("333:",valor3)
	
}