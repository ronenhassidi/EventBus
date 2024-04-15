package queue_util

import (
	"errors"
	"sync"

	log "github.com/gilbarco-ai/event-bus/common/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPool struct {
	connections chan *amqp.Connection
	mu          sync.Mutex
}

func NewRabbitMQPool(url string, maxConnections int) (*RabbitMQPool, error) {
	pool := &RabbitMQPool{
		connections: make(chan *amqp.Connection, maxConnections),
	}

	for i := 0; i < maxConnections; i++ {
		conn, err := amqp.Dial(url)
		if err != nil {
			return nil, err
		}

		pool.connections <- conn
	}

	return pool, nil
}

func (p *RabbitMQPool) GetConnection() (*amqp.Connection, error) {
	select {
	case conn := <-p.connections:
		return conn, nil
	default:
		return nil, errors.New("connection pool is empty")
	}
}

func (p *RabbitMQPool) ReleaseConnection(conn *amqp.Connection) {
	p.connections <- conn
}

func Newchannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Error(err)
		panic(err)
	}

	return ch
}

/*
func SendMessage(ch *amqp.Channel, msg string) {
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully published message to queue")
}
*/

/*
func main() {
	pool, err := NewRabbitMQPool("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pool.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer pool.ReleaseConnection(conn)

	// Use the connection for your RabbitMQ operations here
}
*/
