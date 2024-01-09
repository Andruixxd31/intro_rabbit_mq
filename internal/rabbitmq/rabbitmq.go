package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn *amqp.Connection
}

func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to rabbitmq")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Succesfullly connected to rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
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

	fmt.Println(q)

	err = ch.PublishWithContext(
		context.Background(),
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)

	return nil
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
