package messaging

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	rabbitMQURL string
	channel     *amqp.Channel
}

func NewRabbitMQPublisher(rabbitMQURL string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Failed to open a channel: %v", err)
	}

	return &RabbitMQPublisher{
		rabbitMQURL: rabbitMQURL,
		channel:     ch,
	}, nil
}

func (p *RabbitMQPublisher) Close() {
	p.channel.Close()
}

// Publisher ke RabbitMQ
func (p *RabbitMQPublisher) Publish(key string, message string) error {
	err := p.channel.Publish(
		"",    // Nama exchange
		key,   // Routing key
		false, // Mandatory
		false, // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to publish to RabbitMQ: %v", err)
	}

	fmt.Printf("Published to RabbitMQ: %s\n", message)
	return nil
}
