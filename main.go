package main

import (
	"fmt"
	"github.com/andruiixd/rabbit-mq-intro/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go Rabbitmq intro")
	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}
	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up application")
		fmt.Println(err)
	}

}
