package queue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/cache"
	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/config"
	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/handlers"
	"github.com/casperfj/bachelor/pkg/transaction"
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

func (q *Queue) SubscribeToTransaction(handlers *handlers.Handlers, cache *cache.Cache, configuration *config.Configuration) error {
	// Create go routine for consuming messages
	forever := make(chan bool)

	go func() {
		for d := range q.Messages {
			// Unmarshal transaction
			transaction := &transaction.Transaction{}
			if err := json.Unmarshal(d.Body, transaction); err != nil {
				log.Printf("failed to unmarshal transaction. body: %s, error: %s", d.Body, err.Error())
				d.Ack(false)
				continue
			}

			accountIDs := []uuid.UUID{transaction.SenderAccountID, transaction.ReceiverAccountID}

			// Check if eather account is in other transaction
			// TODO: Handle error
			if isBlocked, err := cache.IsInTransaction(accountIDs); isBlocked || err != nil {
				log.Printf("account is in other transaction. transaction: %s", transaction.ID)
				// Requeue message
				d.Nack(false, true)
				continue
			}

			// Add account ids to cache
			err := cache.BlockAccountIDs(accountIDs)
			if err != nil {
				log.Printf("failed to add account ids to cache. transaction: %s, error: %s", transaction.ID, err.Error())
				// Requeue message
				d.Nack(false, true)
				continue
			}

			// Process transaction
			handlers.ProcessTransaction(transaction, configuration)

			// Remove account ids from cache
			cache.ReleaseAccountIDs(accountIDs)

			// TODO: Catch and handle processing error and notify sender that transaction failed

			// Acknowledge message
			log.Printf("successfully processed transaction. transaction: %s", transaction.ID)
			d.Ack(false)
		}
	}()

	log.Printf("started consuming messages from queue: %s", q.queue.Name)
	<-forever

	// Return because there is no error
	return nil
}
