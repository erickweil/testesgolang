package beginner

// https://www.youtube.com/watch?v=pAXp6o-zWS4
import (
	"fmt"

	"github.com/streadway/amqp"
)

// Inicia a conexão em localhost com usuário guest e senha guest
func getConnection() *amqp.Connection {
	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return conn
}

func getChannel(conn *amqp.Connection) *amqp.Channel {
	ch,err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return ch
}

func declareQueue(ch *amqp.Channel,name string) *amqp.Queue {
	q,err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &q
}

func publishOnQueue(ch *amqp.Channel, queue string, text string) {
	err := ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(text),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Test() {
	fmt.Println("Testando a Conexão Rabbitmq")

	conn := getConnection()
	defer conn.Close()

	ch := getChannel(conn)
	defer ch.Close()

	declareQueue(ch,"teste-queue")
	publishOnQueue(ch,"teste-queue","Primeira Mensagem")


	fmt.Println("Conectado com sucesso!")
}