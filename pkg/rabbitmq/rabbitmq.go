package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ is a wrapper around the amqp091-go package

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ instance

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{}
}

// Connect connects to RabbitMQ

func (r *RabbitMQ) Connect() error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return err
	}
	r.Conn = conn
	return nil
}

// Close closes the connection to RabbitMQ

func (r *RabbitMQ) Close() error {
	return r.Conn.Close()
}

// CreateChannel creates a new channel

func (r *RabbitMQ) CreateChannel() error {
	ch, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	r.Channel = ch
	return nil
}

// CloseChannel closes the channel

func (r *RabbitMQ) CloseChannel() error {
	return r.Channel.Close()
}

// Consume consumes a message from a queue

func (r *RabbitMQ) Consume(out chan amqp.Delivery) error {
	msgs, err := r.Channel.Consume(
		"cookies",
		"go-consumer-docker",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	for d := range msgs {
		fmt.Println("Recebendo mensagem...: ", string(d.Body))
		out <- d
	}
	return nil
}
