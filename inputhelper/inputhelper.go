package inputhelper

import (
	"bufio"
	"os"
)

var scanner *bufio.Scanner

func Readline() string {
	scanner.Scan()
	return scanner.Text()
}

func InputInit() {
	scanner = bufio.NewScanner(os.Stdin)
}
