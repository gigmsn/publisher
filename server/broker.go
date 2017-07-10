package server

import (
	"fmt"

	"github.com/streadway/amqp"
)

// broker wraps the necessary
// information to deal with rabbitmq
type broker struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// create a new broker
func new(addr, name string) (*broker, error) {

	// connect over tpc using plain auth
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to queue: %s", err)
	}

	// open unique connection to send messages to queue
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("could not create queue channel: %s", err)
	}

	// declare a queue to hold messages
	q, err := ch.QueueDeclare(
		name,  // queue name
		false, // durable
		false, // autodelete
		false, // exclusive
		false, // nowait
		nil,   // args amqp.Table
	)
	if err != nil {
		return nil, fmt.Errorf("could not declare queue: %s", err)
	}
	return &broker{conn, ch, q}, nil
}

// publish message to queue
func (b *broker) publish(msgCh chan []byte, doneCh chan bool) {
	for {
		select {
		case msg := <-msgCh:
			b.channel.Publish(
				"",           // exchange string
				b.queue.Name, // queue name
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					Body: msg, // amqp.Publishing
				},
			)
		case <-doneCh:
			close(msgCh)
			<-doneCh
		}
	}
}

// close queue connection and channel
func (b *broker) close() error {
	if err := b.conn.Close(); err != nil {
		return fmt.Errorf("could not close broker connection: %s", err)
	}
	if err := b.channel.Close(); err != nil {
		return fmt.Errorf("could not close broker channel: %s", err)
	}
	return nil
}
