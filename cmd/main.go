package main

import (
	"github.com/daviamorim29/gorabbit/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	instance := rabbitmq.NewRabbitMQ()
	instance.Connect()
	instance.CreateChannel()
	out := make(chan amqp091.Delivery)

	instance.Consume(out)

	for msg := range out {
		println(string(msg.Body))
		msg.Ack(false)
	}
}
