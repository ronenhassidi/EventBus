package configuration

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	log "github.com/gilbarco-ai/event-bus/common/log"
	queueUtil "github.com/gilbarco-ai/event-bus/common/queue/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
)

var RabbitList map[string]Rabbit

type Rabbit struct {
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
	Pool       *queueUtil.RabbitMQPool
	Queue      QueueTemplate
}

type QueueConfig struct {
	pool       *queueUtil.RabbitMQPool
	connection InputConnection
	queueName  string
}

type InputQueueConfig struct {
	Templates      InputTemplate           `yaml:"templates" json:"templates"`
	Exchanges      []InputExchangeTemplate `yaml:"exchanges" json:"exchanges"`
	Queues         []InputQueueTemplate    `yaml:"queues" json:"queues"`
	Consumers      []InputConsumerTemplate `yaml:"consumers" json:"consumers"`
	QueueConsumers []InputQueueConsumers   `yaml:"queueConsumers" json:"queueConsumers"`
	ExchangeQueues []InputExchangeQueues   `yaml:"exchangeQueues" json:"exchangeQueues"`
	Connections    []InputConnection       `yaml:"connections" json:"connections"`
}

type ArgsTable struct {
	Args *amqp.Table
}

type InputTemplate struct {
	Exchanges []ExchangeTemplate `yaml:"exchanges" json:"exchanges"`
	Queues    []QueueTemplate    `yaml:"queues" json:"queues"`
	Consumers []ConsumerTemplate `yaml:"consumers" json:"consumers"`
}

type ExchangeTemplate struct {
	Name       string `yaml:"name" json:"name"`
	Type       string `yaml:"type" json:"type"`
	Durable    bool   `yaml:"durable" json:"durable"`
	AutoDelete bool   `yaml:"autoDelete" json:"autoDelete"`
	Internal   bool   `yaml:"internal" json:"internal"`
	NoWait     bool   `yaml:"noWait" json:"noWait"`
	Args       string `yaml:"args" json:"args"`
}

type QueueTemplate struct {
	Name       string `yaml:"name" json:"name"`
	Durable    bool   `yaml:"durable" json:"durable"`
	AutoDelete bool   `yaml:"autoDelete" json:"autoDelete"`
	Exclusive  bool   `yaml:"exclusive" json:"exclusive"`
	NoWait     bool   `yaml:"noWait" json:"noWait"`
	Args       string `yaml:"args" json:"args"`
}

type ConsumerTemplate struct {
	Name      string `yaml:"name" json:"name"`
	AutoAck   bool   `yaml:"autoAck" json:"autoAck"`
	Exclusive bool   `yaml:"exclusive" json:"exclusive"`
	NoLocal   bool   `yaml:"noLocal" json:"noLocal"`
	NoWait    bool   `yaml:"noWait" json:"noWait"`
	Args      string `yaml:"args" json:"args"`
}

type InputConnection struct {
	Name         string         `yaml:"name" json:"name"`
	ConnectionId string         `yaml:"connectionId" json:"connectionId"`
	ServerURL    string         `yaml:"serverURL" json:"serverURL"`
	User         string         `yaml:"user" json:"user"`
	Password     string         `yaml:"password" json:"password"`
	Port         string         `yaml:"port" json:"port"`
	UIPort       string         `yaml:"UIPort" json:"UIPort"`
	Pool         string         `yaml:"pool" json:"pool"`
	Channels     []InputChannel `yaml:"channels" json:"channels"`
}

type InputChannel struct {
	Name      string          `yaml:"name" json:"name"`
	Instance  int             `yaml:"instance" json:"instance"`
	Exchanges []InputExchange `yaml:"exchanges" json:"exchanges"`
}

/*
type InputConnections struct {
	Connections []InputConnection `yaml:"connections"`
}
*/

type InputQueueTemplate struct {
	Name     string `yaml:"name" json:"name"`
	Template string `yaml:"template" json:"template"`
}

type InputConsumerTemplate struct {
	Name     string `yaml:"name" json:"name"`
	Template string `yaml:"template" json:"template"`
}

type InputQueueConsumers struct {
	Name      string          `yaml:"name" json:"name"`
	Consumers []InputConsumer `yaml:"consumers" json:"consumers"`
}

type InputExchangeQueues struct {
	Name   string        `yaml:"name" json:"name"`
	Queues []InputQueues `yaml:"queues" json:"queues"`
}

type InputConsumer struct {
	Name string `yaml:"name" json:"name"`
}

type InputQueues struct {
	Name string `yaml:"name" json:"name"`
	//	Consumers []InputConsumer `yaml:"consumers" json:"consumers"`
}

type InputQueue struct {
	Name string `yaml:"name" json:"name"`
}

type InputExcahnges struct {
	Name   string       `yaml:"name" json:"name"`
	Queues []InputQueue `yaml:"queues" json:"queues"`
}

type InputExchangeTemplate struct {
	Name     string `yaml:"name" json:"name"`
	Template string `yaml:"template" json:"template"`
}

type InputExchange struct {
	Name string `yaml:"name" json:"name"`
}

type ExchangeQueues struct {
	Name   string          `yaml:"name" json:"name"`
	Queues []QueueTemplate `yaml:"queues" json:"queues"`
}

type ExchangeQueuesTemplate struct {
	Exchange ExchangeTemplate `yaml:"exchange" json:"exchange"`
	Queues   []QueueTemplate  `yaml:"queues" json:"queues"`
}

// Create a new config instance.
var (
	inputQueueConfig *InputQueueConfig
)

func getServerUrl(url string, user string, password string, port string) string {

	serverUrl := strings.Replace(url, "{{user}}", user, 1)
	serverUrl = strings.Replace(serverUrl, "{{password}}", password, 1)
	serverUrl = strings.Replace(serverUrl, "{{port}}", port, 1)
	return serverUrl
}

func getRabbitMqArgs(args string) amqp091.Table {
	table := make(amqp091.Table)
	//args = "'data':[{'x-message-ttl': '60000','x-max-length': '100'}]"
	//	args = `{"x-message-ttl":"60000", "x-max-length":"100"}`
	keyValueMap := map[string]string{}
	if args != "nil" && args != "" {
		err := json.Unmarshal([]byte(args), &keyValueMap)
		if err != nil {
			log.Logger.Error(err)
			panic(err)
		}
		for k, v := range keyValueMap {
			table[k] = v
		}
	}

	return table
}

//================================   Exchange   ==================================================

func getExchangeTemplates(inputQueueConfig *InputQueueConfig) map[string]*ExchangeTemplate {

	exchanges := make(map[string]*ExchangeTemplate)
	for _, t := range inputQueueConfig.Templates.Exchanges {
		excahngeTemplate := ExchangeTemplate{
			Name:       t.Name,
			Type:       t.Type,
			Durable:    t.Durable,
			AutoDelete: t.AutoDelete,
			Internal:   t.Internal,
			NoWait:     t.NoWait,
			Args:       t.Args,
		}
		exchanges[t.Name] = &excahngeTemplate
	}

	return exchanges

}

func getInputExchangeTemplate(exchangeName string, inputQueueConfig *InputQueueConfig) *ExchangeTemplate {

	exchangeTemplate := &ExchangeTemplate{}
	exchangeTemplates := getExchangeTemplates(inputQueueConfig)
	for _, e := range inputQueueConfig.Exchanges {
		if e.Name == exchangeName {
			exchangeTemplate = getExchangeInTemplates(e.Template, exchangeTemplates)
			if exchangeTemplate == nil {
				log.Logger.Error("Invalid exchange name")
				panic("Invalid exchange name")
			}
			return exchangeTemplate
		}

	}

	return exchangeTemplate

}

func getExchanges(inputQueueConfig *InputQueueConfig) []InputExchangeTemplate {

	exchangeList := []InputExchangeTemplate{}
	for _, e := range inputQueueConfig.Exchanges {
		inputExchangeTemplate := InputExchangeTemplate{
			Name:     e.Name,
			Template: e.Template,
		}
		exchangeList = append(exchangeList, inputExchangeTemplate)
	}
	return exchangeList
}

func getExchangeInTemplates(exchangeName string, exchangeTemplates map[string]*ExchangeTemplate) *ExchangeTemplate {
	exchangeTemplate := &ExchangeTemplate{}
	for _, e := range exchangeTemplates {
		if e.Name == exchangeName {
			exchangeTemplate = &ExchangeTemplate{
				Name:       e.Name,
				Type:       e.Type,
				Durable:    e.Durable,
				AutoDelete: e.AutoDelete,
				Internal:   e.Internal,
				NoWait:     e.NoWait,
				Args:       e.Args,
			}
			return exchangeTemplate
		}
	}
	return nil
}

func getExchangeTemplateInExchanges(exchangeName string, exchangeTemplates map[string]*ExchangeTemplate) *ExchangeTemplate {
	exchangeTemplate := &ExchangeTemplate{}
	for _, e := range exchangeTemplates {
		if e.Name == exchangeName {
			exchangeTemplate = &ExchangeTemplate{
				Name:       e.Name,
				Type:       e.Type,
				Durable:    e.Durable,
				AutoDelete: e.AutoDelete,
				Internal:   e.Internal,
				NoWait:     e.NoWait,
				Args:       e.Args,
			}
			return exchangeTemplate
		}
	}
	return nil
}

func getExchangeTemplate(exchanges map[string]*ExchangeQueuesTemplate, exchangeName string) ExchangeTemplate {

	var exchange ExchangeTemplate
	for _, e := range exchanges {
		if e.Exchange.Name == exchangeName {
			exchange = e.Exchange
			break
		}
	}
	return exchange
}

func getExchangeQueueTemplates(exchanges map[string]*ExchangeQueuesTemplate, exchangeName string) []QueueTemplate {

	var queues []QueueTemplate
	for _, e := range exchanges {
		if e.Exchange.Name == exchangeName {
			queues = e.Queues
			break
		}
	}
	return queues
}

//================================   Queue   ==================================================

func getQueueTemplates(inputQueueConfig *InputQueueConfig) map[string]*QueueTemplate {

	queues := make(map[string]*QueueTemplate)
	for _, t := range inputQueueConfig.Templates.Queues {
		queueTemplate := QueueTemplate{
			Name:       t.Name,
			Durable:    t.Durable,
			AutoDelete: t.AutoDelete,
			Exclusive:  t.Exclusive,
			NoWait:     t.NoWait,
			Args:       t.Args,
		}
		queues[t.Name] = &queueTemplate
	}

	return queues

}

func getQueuess(inputQueueConfig *InputQueueConfig) []InputQueueTemplate {

	queueList := []InputQueueTemplate{}
	for _, q := range inputQueueConfig.Queues {
		inputQueueTemplate := InputQueueTemplate{
			Name:     q.Name,
			Template: q.Template,
		}
		queueList = append(queueList, inputQueueTemplate)
	}
	return queueList
}

func getQueueInTemplates(queueName string, queueTemplates map[string]*QueueTemplate) *QueueTemplate {
	queueTemplate := &QueueTemplate{}
	for _, q := range queueTemplates {
		if q.Name == queueName {
			queueTemplate = &QueueTemplate{
				Name:       q.Name,
				Durable:    q.Durable,
				AutoDelete: q.AutoDelete,
				Exclusive:  q.Exclusive,
				NoWait:     q.NoWait,
				Args:       q.Args,
			}
			return queueTemplate
		}
	}
	return nil
}

func getQueueTemplate(queueName string, inputQueueConfig *InputQueueConfig) *QueueTemplate {

	queueTemplate := &QueueTemplate{}
	queueTemplates := getQueueTemplates(inputQueueConfig)
	for _, e := range inputQueueConfig.Queues {
		if e.Name == queueName {
			queueTemplate = getQueueInTemplates(e.Template, queueTemplates)
			if queueTemplate == nil {
				log.Logger.Error("Invalid queue name")
				panic("Invalid queue name")
			}
			return queueTemplate
		}

	}

	return queueTemplate

}

//================================   Consumer   ==================================================

func getConsumerTemplates(inputQueueConfig *InputQueueConfig) map[string]*ConsumerTemplate {

	consumers := make(map[string]*ConsumerTemplate)
	for _, t := range inputQueueConfig.Templates.Consumers {
		inputConsumerTemplate := ConsumerTemplate{
			Name:      t.Name,
			AutoAck:   t.AutoAck,
			Exclusive: t.Exclusive,
			NoLocal:   t.NoLocal,
			NoWait:    t.NoWait,
			Args:      t.Args,
		}
		consumers[t.Name] = &inputConsumerTemplate
	}

	return consumers

}

func getConsumers(inputQueueConfig *InputQueueConfig) []InputConsumerTemplate {

	consumerList := []InputConsumerTemplate{}
	for _, c := range inputQueueConfig.Exchanges {
		inputConsumerTemplate := InputConsumerTemplate{
			Name:     c.Name,
			Template: c.Template,
		}
		consumerList = append(consumerList, inputConsumerTemplate)
	}
	return consumerList
}

//================================   Exchange queue   ==================================================

func getExchangeQueues(inputQueueConfig *InputQueueConfig) map[string]*ExchangeQueuesTemplate {
	//	exchangeTemplates := getExchangeTemplates(inputQueueConfig)
	//queueTemplates := getQueueTemplates(inputQueueConfig)
	exchangeQueues := make(map[string]*ExchangeQueuesTemplate)
	exchangeTemplate := &ExchangeTemplate{}
	queueTemplate := &QueueTemplate{}
	for _, eq := range inputQueueConfig.ExchangeQueues {
		if eq.Name == "" {
			log.Logger.Error("Invalid exchange name")
			panic("Invalid exchange name")
		}
		exchangeTemplate = getInputExchangeTemplate(eq.Name, inputQueueConfig)
		if exchangeTemplate == nil {
			log.Logger.Error("Invalid exchange name")
			panic("Invalid exchange name")
		}
		exchangeTemplate.Name = eq.Name
		queueList := []QueueTemplate{}
		for _, qs := range eq.Queues {
			if qs.Name == "" {
				log.Logger.Error("Invalid queue name")
				panic("Invalid queue name")
			}
			queueTemplate = getQueueTemplate(qs.Name, inputQueueConfig)
			if queueTemplate == nil {
				log.Logger.Error("Invalid queue name")
				panic("Invalid queue name")
			}
			queueTemplate.Name = qs.Name
			queueList = append(queueList, *queueTemplate)
		}
		exchangeQueuesTemplate := ExchangeQueuesTemplate{
			Exchange: *exchangeTemplate,
			Queues:   queueList,
		}
		exchangeQueues[eq.Name] = &exchangeQueuesTemplate
	}

	return exchangeQueues

}

func GetExchanges(inputQueueConfig *InputQueueConfig) map[string]*ExchangeQueuesTemplate {
	exchanges := make(map[string]*ExchangeQueuesTemplate)
	exchanges = getExchangeQueues(inputQueueConfig)
	return exchanges
}

func CreateQueueList(inputQueueConfig *InputQueueConfig) map[string]Rabbit {

	exchanges := GetExchanges(inputQueueConfig)
	//var rabbitList []Rabbit = make([]Rabbit, 0)
	queueMap := make(map[string]Rabbit)
	for _, c := range inputQueueConfig.Connections {
		serverUrl := getServerUrl(c.ServerURL,
			c.User,
			c.Password,
			c.Port)
		//	log.Logger.Info(serverUrl)
		// Get pool size
		poolSize, err := strconv.Atoi(c.Pool)
		if err != nil {
			log.Logger.Error(err)
			panic(err)
		}

		// create pool
		pool, err := queueUtil.NewRabbitMQPool(serverUrl, poolSize)
		if err != nil {
			log.Logger.Error(err)
			panic(err)
		}
		for _, c := range c.Channels {

			conn, err := pool.GetConnection()
			if err != nil {
				log.Logger.Println(err)
				panic(err)
			}
			fmt.Println("Succesfully connected to RabbitMQ")
			ch, err := conn.Channel()
			if err != nil {
				log.Logger.Println(err)
				panic(err)
			}
			fmt.Println("Succesfully get RabbitMQ channel")
			for _, e := range c.Exchanges {
				exchangeTemplate := getExchangeTemplate(exchanges, e.Name)
				err = ch.ExchangeDeclare(
					exchangeTemplate.Name,                  // exchange name
					exchangeTemplate.Type,                  // exchange type
					exchangeTemplate.Durable,               // durable
					exchangeTemplate.AutoDelete,            // auto-deleted
					exchangeTemplate.Internal,              // internal
					exchangeTemplate.NoWait,                // no-wait
					getRabbitMqArgs(exchangeTemplate.Args), // arguments
				)
				if err != nil {
					log.Logger.Println(err)
					panic(err)
				}

				qs := getExchangeQueueTemplates(exchanges, e.Name)
				for _, q := range qs {
					_, err := ch.QueueDeclare(
						q.Name,
						q.Durable,
						q.AutoDelete,
						q.Exclusive,
						q.NoWait,
						getRabbitMqArgs(q.Args),
					)
					if err != nil {
						log.Logger.Println(err)
						panic(err)
					}
					rabbit := Rabbit{
						Pool:       pool,
						Connection: conn,
						Channel:    ch,
						Queue:      q,
					}
					queueMap[rabbit.Queue.Name] = rabbit
				}

			}
		}
	}
	return queueMap
}

/*
func createRabbitQueue(pool *queueUtil.RabbitMQPool, q QueueTemplate) Rabbit {

	if pool == nil {
		log.Logger.Println("Invalid pool")
		panic("Invalid pool")
	}

	conn, err := pool.GetConnection()
	if err != nil {
		log.Logger.Println(err)
		panic(err)
	}
	fmt.Println("Succesfully connected to RabbitMQ")
	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Println(err)
		panic(err)
	}
	fmt.Println("Succesfully get RabbitMQ channel")

	queue, err := ch.QueueDeclare(
		q.Name,
		q.Durable,
		q.AutoDelete,
		q.Exclusive,
		q.NoWait,
		getRabbitMqArgs(q.Args),
	)

	if err != nil {
		log.Logger.Println(err)
		panic(err)
	}
	fmt.Println(queue)

	rabbit := Rabbit{
		Pool:       pool,
		Connection: conn,
		Channel:    ch,
		Queue:      q,
	}

	return rabbit
}*/

//---------------------------    public   ---------------------------------------------------
/*
func FillQueueTemplateValues(inputQueueConfig *InputQueueConfig) *InputQueueConfig {
	exchangeTemplates := getExchangeTemplates(inputQueueConfig)
	queueTemplates := getQueueTemplates(inputQueueConfig)
	consumerTemplates := getConsumerTemplates(inputQueueConfig)
	for _, con := range inputQueueConfig.Connections {
		for _, c := range con.Channels {
			if c.Name == "" {
				log.Logger.Error("Invalid connection channel name")
				panic("Invalid connection channel name")
			}
			if c.Instance <= 0 {
				log.Logger.Error("Invalid number of instances")
				panic("Invalid number of instances")
			}
			if len(c.Exchanges) > 0 {
				queueTemplate, ok := queueTemplates[q.Template]
				if ok {
					q.Durable = queueTemplate.Durable
					q.AutoDelete = queueTemplate.AutoDelete
					q.Exclusive = queueTemplate.Exclusive
					q.NoWait = queueTemplate.NoWait
					q.Args = queueTemplate.Args
				}
			}
			if len(q.Consumers) > 0 {
				for _, c := range q.Consumers {
					if len(c.Template) > 0 {
						consumerTemplate, ok := consumerTemplates[c.Template]
						if ok {
							c.AutoAck = consumerTemplate.AutoAck
							c.Exclusive = consumerTemplate.Exclusive
							c.NoLocal = consumerTemplate.NoLocal
							c.NoWait = consumerTemplate.NoWait
							c.Args = consumerTemplate.Args
						}
					}
				}
			}
		}
	}
	return inputQueueConfig
}
*/

/*
func CreateQueueList(inputQueueConfig *InputQueueConfig) map[string]Rabbit {

	//var rabbitList []Rabbit = make([]Rabbit, 0)
	queueMap := make(map[string]Rabbit)
	for _, c := range inputQueueConfig.Connections {
		serverUrl := getServerUrl(c.ServerURL,
			c.User,
			c.Password,
			c.Port)
		//	log.Logger.Info(serverUrl)

		// Get pool size
		poolSize, err := strconv.Atoi(c.Pool)
		if err != nil {
			log.Logger.Error(err)
			panic(err)
		}

		// create pool
		pool, err := queueUtil.NewRabbitMQPool(serverUrl, poolSize)
		if err != nil {
			log.Logger.Error(err)
			panic(err)
		}
		for _, c := range c.Channels {
			exchangeTemplate := getExchange(inputQueueConfig, c.Name)
			err = ch.ExchangeDeclare(
				"direct_logs", // name
				"direct",      // type
				true,          // durable
				false,         // auto-deleted
				false,         // internal
				false,         // no-wait
				nil,           // arguments
			)

			for _, q := range c.Queues {

				if strings.Contains(q.Name, "{{i}}") {
					for i := 0; i < q.Multiplier; i++ {
						newQueue := q
						newQueue.Consumers = nil
						newQueue.Consumers = append(newQueue.Consumers, q.Consumers...)
						name := "queue" + strconv.Itoa(i+1)
						newQueue.Name = name
						rabbit := createRabbitQueue(pool, newQueue)
						queueMap[name] = rabbit
					}
				} else {
					rabbit := createRabbitQueue(pool, q)
					queueMap[rabbit.Queue.Name] = rabbit
				}
			}
		}
	}
	return queueMap
}
*/

/*
func createRabbitQueue(pool *queueUtil.RabbitMQPool, q InputQueue) Rabbit {

		if pool == nil {
			logger.Error.Println("Invalid pool")
			panic("Invalid pool")
		}

		conn, err := pool.GetConnection()
		if err != nil {
			logger.Error.Println(err)
			panic(err)
		}
		fmt.Println("Succesfully connected to RabbitMQ")
		ch, err := conn.Channel()
		if err != nil {
			logger.Error.Println(err)
			panic(err)
		}
		fmt.Println("Succesfully get RabbitMQ channel")

		queue, err := ch.QueueDeclare(
			q.Name,
			q.Durable,
			q.AutoDelete,
			q.Exclusive,
			q.NoWait,
			getRabbitMqArgs(q.Args),
		)

		if err != nil {
			logger.Error.Println(err)
			panic(err)
		}
		fmt.Println(queue)

		rabbit := Rabbit{
			Pool:       pool,
			Connection: conn,
			Channel:    ch,
			Queue:      q,
		}

		return rabbit
	}

func GetQueuesConfiguration(inputQueueConfig *InputQueueConfig) map[string]*QueueConfig {

	queues := make(map[string]*QueueConfig)
	for i := 0; i < len(inputQueueConfig.Connections); i++ {
		conn := inputQueueConfig.Connections[i]
		serverUrl := getServerUrl(inputQueueConfig.Connections[i].ServerURL,
			inputQueueConfig.Connections[i].User,
			inputQueueConfig.Connections[i].Password,
			inputQueueConfig.Connections[i].Port)

		fmt.Println(serverUrl)

		// Get pool size
		poolSize, err := strconv.Atoi(conn.Pool)
		if err != nil {
			logger.Error.Fatal(err)
			panic(err)
		}

		// create pool
		pool, err := queueUtil.NewRabbitMQPool(serverUrl, poolSize)
		if err != nil {
			logger.Error.Fatal(err)
			panic(err)
		}

		//for j := 0; j < len(conn.Queues); j++ {
		for _, q := range conn.Queues {
			queueConfig := QueueConfig{
				queueName:  q.Name,
				pool:       pool,
				connection: conn,
			}
			queues[q.Name] = &queueConfig
		}

	}

	return queues

}

	func CreateQueues(queuesConfig map[string]*QueueConfig) {
		n := 0
		keys := make([]string, len(queuesConfig))
		for k := range queuesConfig {
			keys[n] = k
			n++
		}
		for k, v := range queuesConfig {
			conn, err := queuesConfig[k].pool.GetConnection()
			if err != nil {
				logger.Error.Println(err)
				panic(err)
			}
			defer conn.Close()
			log.Println("Succesfully connected to queue")

			ch, err := conn.Channel()
			if err != nil {
				logger.Error.Println(err)
				panic(err)
			}
			defer ch.Close()

			for _, queueInput := range v.connection.Queues {
				q, err := ch.QueueDeclare(
					k,
					queueInput.Durable,
					queueInput.AutoDelete,
					queueInput.Exclusive,
					queueInput.NoWait,
					getRabbitMqArgs(queueInput.Args),
				)
				if err != nil {
					logger.Error.Println(err)
					panic(err)
				}
				logger.Info.Println("Queue created - " + q.Name)
			}

		}

}

	func GetConsumerList(inputQueueConfig *InputQueueConfig) map[string]*InputConsumer {
		consumerMap := make(map[string]*InputConsumer)
		for _, c := range inputQueueConfig.Templates.Consumers {
			inputConsumer := InputConsumer{
				Name:      c.Name,
				AutoAck:   c.AutoAck,
				Exclusive: c.Exclusive,
				NoLocal:   c.NoLocal,
				NoWait:    c.NoWait,
				Args:      c.Args,
			}
			consumerMap[c.Name] = &inputConsumer
		}
		return consumerMap
	}

func SendMsgByChannel(rabbit Rabbit, msg string) {

	err := rabbit.Channel.Publish(
		"",
		rabbit.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		logger.Error.Println(err)
		//	panic(err)
	}
	fmt.Println("Successfully published message to queue - " + rabbit.Queue.Name)

}

func ConsumeMsg(rabbit Rabbit, inputConsumer *InputConsumer, queueName string) {

	conn, err := rabbit.Pool.GetConnection()
	if err != nil {
		logger.Error.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Succesfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		logger.Error.Println(err)
		panic(err)
	}
	defer ch.Close()
	fmt.Println(inputConsumer.AutoAck)
	msgs, err := ch.Consume(
		queueName,
		"",
		inputConsumer.AutoAck,
		inputConsumer.Exclusive,
		inputConsumer.NoLocal,
		inputConsumer.NoWait,
		getRabbitMqArgs(inputConsumer.Args),
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()
	logger.Info.Println("Successfully connected to queue")
	logger.Info.Println(" [*] - Waitint for messages")
	<-forever

}
*/
func SendMsgByChannel(rabbit Rabbit, msg string) {}
