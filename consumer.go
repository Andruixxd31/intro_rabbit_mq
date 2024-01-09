package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Added another main function just to have the consumer and the rabbitMQ service inside the same repo
func main() {
	fmt.Println("Consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received msg: %s\n", d.Body)
		}
	}()

	fmt.Println("Succesfully connected to rabbitmq instance")
	fmt.Println("[*] - waiting for messages")
	<-forever
}
