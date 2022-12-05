package main

import (
	"example/inputhelper"
	"example/rabbitmq/beginner"
	"example/rabbitmq/tutorial1"
	"example/rabbitmq/tutorial2"
	"fmt"
)

func execTutorial1() {
	var channel = "hello"
	fmt.Println("1 - Send, 2 - Receive Server")
	var qual int
	fmt.Scan(&qual)

	if qual == 1 {
		inputhelper.InputInit()

		for {
			fmt.Println("Escreva uma mensagem:")
			var msg = inputhelper.Readline()
			tutorial1.TestSend(channel,msg)
		}
	} else {
		tutorial1.TestReceive(channel)
	}
}

func main() {
	fmt.Println("----------- RABBITMQ ------------")
	fmt.Println("1 - TUTORIAL1: Enviar e Receber")
	fmt.Println("2 - TUTORIAL2: Worker")
	fmt.Println("3 - Testes: Conexão e Testes")
	var qual int
	fmt.Scan(&qual)

	switch qual {
	case 1:
		execTutorial1()
	case 2:
		fmt.Println("TUTORIAL2: Worker")
		fmt.Println("	1 - Worker")
		fmt.Println("	2 - Task")
		fmt.Scan(&qual)
		if qual == 1 {
			tutorial2.MainWorker()
		}
		if qual == 2 {
			tutorial2.MainTask()
		}
	case 3:
		beginner.Test()
	}

}
