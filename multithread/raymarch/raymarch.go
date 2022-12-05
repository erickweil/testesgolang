package raymarch

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

/*
Código que gera uma imagem renderizada com Raymarching

A ideia é que raios partem da câmera e encontram a superfície
de objetos definidos por equações matemáticas. Cada pixel é
um raio que será traçado

As superfícies são definidas como funções estimadoras de
distância, basicamente recebem uma posição e retornam a distância
desta posição até a superfície mais próxima. Se estiver dentro
retorna esta distância negativa. 0 seria EXATAMENTE na superfície

O algoritmo raymarching calcula a distância até a superfície com
essa função e anda na direção do raio a distância calculada, assim
o raio irá andar apenas o quanto tem certeza que não há nada,
chegando cada vez mais perto da superfície sem de fato atravessá-la

Uma vez que o raio tenha encontrado a superfície é calculado a iluminação
neste ponto, utilizando regras de iluminação de lambert, que utiliza
o produto escalar da direção da luz e a normal da superfície.
*/

// Função que será chamada para terminar a aplicação se
// ocorrer um erro (Caso a imagem não exista no disco por exemplo)
func check(err error) {  
    if err != nil {  
        panic(err)  
    }  
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

func renderizar(largura int, altura int) image.Image {
	var img *image.RGBA = image.NewRGBA(image.Rect(0,0,largura,altura))
	
	var alturaf = float64(altura)
	var alturapor2 = alturaf/2.0
	var largurapor2 = float64(largura)/2.0
	// PARA CADA LINHA DA IMAGEM
	for x := 0; x < largura; x++ {
		// PARA CADA CÉLULA NESTA LINHA
		for y := 0; y < altura; y++ {
			// remapear de pixel para fragmento
			// está certo dividir os dois pela altura, é para ficar correto o aspecto
			// e subtrai da altura por 2 porque precisa que o {0.0,0.0} seja o centro
			xf := ((float64(x) - largurapor2) / alturaf) 
			yf := ((float64(y) - alturapor2) / alturaf)

			// chama a função do renderizador
			var cor V3 = renderizarPixel(xf,yf)

			// converte de float 0.0-1.0 para byte 0-255
			var novoPixel = color.RGBA{
				uint8(cor.X*255.0),
				uint8(cor.Y*255.0),
				uint8(cor.Z*255.0),
				uint8(255), // sem transparência
			}

			// salvando o pixel alterado na nova imagem
			img.Set(x,y,novoPixel)
		}
	}
	return img
}

func Main() {
	main()
}

func main() {
	// FULL HD
	start := time.Now() // Medindo tempo

	var img = renderizar(1920,1080)

	elapsed := time.Since(start)
	fmt.Printf("Demorou %s para gerar a imagem\n", elapsed)

	SalvarImagem(img,"./raymarch/render.png")
}
