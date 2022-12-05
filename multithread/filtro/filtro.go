package filtro

/*
Código que lê uma imagem png do disco e aplica um filtro
preto e branco a ela.

A imagem deve estar na mesma pasta deste arquivo, com o nome
parrot.png

A imagem processada será salva com o nome parro-pretoebranco.png
*/

import (
	// Para poder ler a imagem do disco

	"image/color"
	"image/png"
	"os"
	"time"

	// Para manipular a imagem
	"image"

	// Para escrever na tela / lidar com strings
	"fmt"
)

// Função que será chamada para terminar a aplicação se
// ocorrer um erro (Caso a imagem não exista no disco por exemplo)
func check(err error) {  
    if err != nil {  
        panic(err)  
    }  
}

/*
Carregar imagem do disco, Basicamente irá ler o arquivo como qualquer outro
E em seguida irá decodificar o arquivo o interpretando como uma Imagem
E por último retornará um slice 2d com os pixels
fontes:
	https://go-recipes.dev/more-working-with-images-in-go-30b11ab2a9f0
	https://usman.it/image-manipulation-in-golang/
*/
func CarregarImagem(caminho string) image.Image {
	// LENDO ARQUIVO
	arquivo,err := os.Open(caminho) // Abre o arquivo para leitura
	check(err) // verifica se houve algum erro
	defer arquivo.Close() // Quando essa função terminar o arquivo será fechado

	// DECODIFICANDO IMAGEM
	img, formato, err := image.Decode(arquivo)
	check(err) // verifica se houve algum erro
	fmt.Println("Formato da Imagem:",formato)
	
	return img
}

/*
Salva imagem no disco no caminho especificado
*/
func SalvarImagem(img image.Image, caminho string) {
	f,err:= os.Create(caminho)
	check(err)
	defer f.Close()

	err = png.Encode(f, img)
	check(err)
}

/*
Aplica um filtro preto e branco nesta imagem
fonte: https://medium.com/@damithadayananda/image-processing-with-golang-8f20d2d243a2
*/
func FiltroPretoEBranco(img image.Image) image.Image {
	var tamanho image.Point = img.Bounds().Size()
	var novaImg *image.RGBA = image.NewRGBA(image.Rect(0,0,tamanho.X,tamanho.Y))

	// PARA CADA LINHA DA IMAGEM
	for x := 0; x < tamanho.X; x++ {
		// PARA CADA CÉLULA NESTA LINHA
		for y := 0; y < tamanho.Y; y++ {
			// lendo o pixel nas coordenadas x e y respectivas
			var pixel = img.At(x,y)

			// Acessa os valores RGBA indo de 0 a 255 cada um
			var rgba = color.RGBAModel.Convert(pixel).(color.RGBA)

			// Processando filtro preto e branco
			// Basicamente precisa calcular a média dos componentes RGB

			// converte para float, para ter precisão no cálculo da média
			var R = float64(rgba.R)
			var G = float64(rgba.G)
			var B = float64(rgba.B)
		
			// Poderia ter pesos fazendo uma cor importar mais que a outra
			// mas a media deixa cinza de um jeito aceitável.
			var media = (R + G + B) / 3.0

			// converte para uint8 que é o formato que vai de 0 a 255 devolta
			var cor_cinza = uint8(media)

			// cria o pixel com os valores iguais em R, G e B, deixando cinza
			var novoPixel = color.RGBA{
				cor_cinza,
				cor_cinza,
				cor_cinza,
				rgba.A, // continua a mesma trasnparência
			}

			// salvando o pixel alterado na nova imagem
			novaImg.Set(x,y,novoPixel)
		}
	}
	return novaImg
}

func Main() {
	main()
}

// Executar e medir tempo.
func main() {
	var img = CarregarImagem("./filtro/parrot.png")

	start := time.Now() // Medindo tempo

	var novaImg = FiltroPretoEBranco(img)

	elapsed := time.Since(start)
	fmt.Printf("Demorou %s para aplicar o filtro\n", elapsed)

	SalvarImagem(novaImg,"./filtro/parrot-pretoebranco.png")

}