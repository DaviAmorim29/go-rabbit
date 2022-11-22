package rabbit

import amqp "github.com/rabbitmq/amqp091-go"

// RabbitMQ is a wrapper around the amqp091-go package

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ instance

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{}
	// rabbit := &RabbitMQ{}
	// rabbit.Connect()
	// rabbit.CreateChannel()

	// return rabbit
}

// Connect connects to RabbitMQ

func (r *RabbitMQ) Connect() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

// Publish publishes a message to a queue

func (r *RabbitMQ) Publish(queueName string, message string) error {
	q, err := r.Channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = r.Channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Consume consumes a message from a queue

func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	q, err := r.Channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	msgs, err := r.Channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
