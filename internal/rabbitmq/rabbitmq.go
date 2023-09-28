package rabbitmq

import (
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
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Succesfullly connected to rabbitmq")
	return nil
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
