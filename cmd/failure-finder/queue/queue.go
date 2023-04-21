package queue

import (
	"fmt"
	"log"

	"github.com/casperfj/bachelor/cmd/failure-finder/config"
	"github.com/casperfj/bachelor/cmd/failure-finder/handlers"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	queue      *amqp.Queue
	Messages   <-chan amqp.Delivery
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

	// Subscribe to queue
	messages, err := channel.Consume(
		q.Name,
		"",
		false,
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
		Messages:   messages,
	}

	// Return queue
	return queue, nil
}

func (q *Queue) SubscribeToAccounts(handlers *handlers.Handlers, configuration *config.Configuration) error {
	// Create go routine for consuming messages
	forever := make(chan bool)

	go func() {
		for d := range q.Messages {
			// Get account id from message
			accountID, err := uuid.Parse(string(d.Body))
			if err != nil {
				log.Printf("failed to parse account id from message. error: %s", err.Error())
				d.Ack(false)
				continue
			}

			log.Printf("received account id: %s", accountID.String())

			// Process transaction
			handlers.ProcessAccount(accountID)

			// TODO: Catch and handle processing error

			// Acknowledge message
			log.Printf("successfully processed account. account id: %s", accountID.String())
			d.Ack(false)
		}
	}()

	log.Printf("started consuming messages from queue: %s", q.queue.Name)
	<-forever

	// Return because there is no error
	return nil
}
