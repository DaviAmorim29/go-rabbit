package main

import (
	"fmt"

	"github.com/daviamorim29/gorabbit/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	instance := rabbitmq.NewRabbitMQ()
	err := instance.Connect()
	if err != nil {
		panic(err)
	}
	errChan := instance.CreateChannel()
	if errChan != nil {
		panic(errChan)
	}
	out := make(chan amqp091.Delivery)

	go instance.Consume(out)
	fmt.Println("Consumindo mensagens...")
	for msg := range out {
		println(string(msg.Body))
		msg.Ack(false)
	}
}
