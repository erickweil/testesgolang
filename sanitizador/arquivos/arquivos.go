package arquivos

import (
	"fmt"
	"io/fs"
	"os"

	//"github.com/k3a/html2text"
	"github.com/pkg/errors"
	"jaytaylor.com/html2text"
)

func validateDirectory(dirPath string) (string, error) {

    info, err := os.Stat(dirPath)
    if os.IsNotExist(err) {
        return "", errors.Wrapf(err, "Directory does not exist: %v\n", dirPath)
    }
    if err != nil {
        return "", errors.Wrapf(err, "Directory error: %v\n", dirPath)

    }
    if !info.IsDir() {
        return "", errors.Errorf("Directory is a file, not a directory: %#v\n", dirPath)
    }
    return dirPath, nil
}

func validateFile(filePath string) (string, error) {

    info, err := os.Stat(filePath)
    if os.IsNotExist(err) {
        return "", errors.Wrapf(err, "File does not exist: %v\n", filePath)
    }
    if err != nil {
        return "", errors.Wrapf(err, "File error: %v\n", filePath)

    }
    if info.IsDir() {
        return "", errors.Errorf("File is a Directory, not a file: %#v\n", filePath)
    }
    return filePath, nil
}

func ListarArquivos(caminho string) ([]fs.DirEntry, error) {

	caminho, err := validateDirectory(caminho)
	if err != nil { return nil,fmt.Errorf("Caminho inválido: %w",err) }

	files, err := os.ReadDir(caminho)
	if err != nil { return nil,fmt.Errorf("Não foi possível listar: %w",err) }

	return files, nil
}

func LerArquivo(caminho string) ([]byte, error) {
	
	caminho, err := validateFile(caminho)
	if err != nil { return nil,fmt.Errorf("Caminho inválido: %w",err) }

	binario, err := os.ReadFile(caminho)
	if err != nil { return nil,fmt.Errorf("Erro na leitura: %w",err) }

	return binario,nil
}

func ConverterTexto(b []byte) string {
	return string(b[:])
}

func RemoverTagsHTML(html string) (string,error) {
	return html2text.FromString(html, html2text.Options{PrettyTables: false})
}

func SalvarArquivo(caminho string, conteudo string) error {
	return os.WriteFile(caminho,[]byte(conteudo), 0644)
}