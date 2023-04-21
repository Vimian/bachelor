package queue

import (
	"fmt"
	"log"

	"github.com/casperfj/bachelor/cmd/failure-finder-manager/config"
	"github.com/streadway/amqp"
)

type Queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	queue      *amqp.Queue
}

func NewQueue(configuration *config.Configuration) (*Queue, error) {
	// Connect to rabbitmq
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", configuration.RabbitMQ.User, configuration.RabbitMQ.Password, configuration.RabbitMQ.Host, configuration.RabbitMQ.Port))
	if err != nil {
		log.Printf("failed to connect to rabbitmq. error: %s", err.Error())
		return nil, err
	}

	// Open channel
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declare queue
	q, err := channel.QueueDeclare(
		configuration.RabbitMQ.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Initialize queue
	queue := &Queue{
		Connection: conn,
		Channel:    channel,
		queue:      &q,
	}

	// Return queue
	return queue, nil
}

func (q *Queue) PublishAccountID(accountID string) error {
	// Publish account id
	err := q.Channel.Publish(
		"",
		q.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "string",
			Body:        []byte(accountID),
		},
	)
	if err != nil {
		return err
	}

	// Return because there is no error
	return nil
}
