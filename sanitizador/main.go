package main

import (
	"example/sanitizador/arquivos"
	"fmt"
	"os"
	"strings"

	// Go provides a `flag` package supporting basic
	// command-line flag parsing. We'll use this package to
	// implement our example command-line program.
	"flag"
)

var SEP = string(os.PathSeparator)

func runSanitize(dir string, dest string) {
	fmt.Println("Sanitizando",dir,"...")

	dirls, err := arquivos.ListarArquivos(dir)
	if err != nil { fmt.Println("Erro!", err) }

	os.Mkdir(dest,os.ModePerm)

	for _,d := range dirls {
		fmt.Println("Nome:",d.Name())
		if !d.IsDir() {
			var caminho = dir+SEP+d.Name()
			var conteudo,err = arquivos.LerArquivo(caminho)
			if err != nil { fmt.Println("Erro!", err); continue; }
			
			var conteudotxt = arquivos.ConverterTexto(conteudo)
			if strings.HasSuffix(d.Name(),".html") {
				conteudotxt,err = arquivos.RemoverTagsHTML(conteudotxt)
				if err != nil { fmt.Println("Erro!", err); continue; }
			}
			
			
			var caminho_dest = dest+SEP+d.Name()
			if !strings.HasSuffix(d.Name(),".txt") {
				caminho_dest = caminho_dest+".txt"
			}
			arquivos.SalvarArquivo(caminho_dest,conteudotxt)
		}
	}

	
}

func main() {

	//recursivo := flag.Bool("r",false,"Pesquisa recursiva")
	sanitize := flag.Bool("sanitize",false,"Executar Sanitização, gerar .txt a partir dos HTML e etc...")
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		dir = "./teste"
	}

	if *sanitize {
		runSanitize(dir,dir+"-sanitized")
	} else {
		fmt.Println("Nenhuma ação definida")
	}
}