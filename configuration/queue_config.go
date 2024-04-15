package configuration

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gilbarco-ai/event-bus/common/constants"
	log "github.com/gilbarco-ai/event-bus/common/log"
	queueUtil "github.com/gilbarco-ai/event-bus/common/queue/rabbitmq"
	utils "github.com/gilbarco-ai/event-bus/common/utils"
	"github.com/streadway/amqp"
)

var RabbitList map[string]Rabbit

type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Pool       *queueUtil.RabbitMQPool
	Queue      InputQueue
}

type QueueConfig struct {
	pool       *queueUtil.RabbitMQPool
	connection InputConnection
	queueName  string
}

type InputQueueConfig struct {
	Templates      InputTemplate           `yaml:"templates,mapstructure" json:"templates,mapstructure"`
	Exchanges      []InputExchangeTemplate `yaml:"exchanges,mapstructure" json:"exchanges,mapstructure"`
	Queues         []InputQueueTemplate    `yaml:"queues,mapstructure" json:"queues,mapstructure"`
	Consumers      []InputConsumerTemplate `yaml:"consumers,mapstructure" json:"consumers,mapstructure"`
	QueueConsumers []InputQueueConsumers   `yaml:"queueConsumers,mapstructure" json:"queueConsumers,mapstructure"`
	ExchangeQueues []InputExchangeQueues   `yaml:"exchangeQueues,mapstructure" json:"exchangeQueues,mapstructure"`
	Connections    []InputConnection       `yaml:"connections,mapstructure" json:"connections,mapstructure"`
}

type ArgsTable struct {
	Args *amqp.Table
}

type InputTemplate struct {
	Exchanges []ExchangeTemplate `yaml:"exchanges,mapstructure" json:"exchanges,mapstructure"`
	Queues    []QueueTemplate    `yaml:"queues,mapstructure" json:"queues,mapstructure"`
	Consumers []ConsumerTemplate `yaml:"consumers,mapstructure" json:"consumers,mapstructure"`
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
	Channels     []InputChannel `yaml:"channels,mapstructure" json:"channels,mapstructure"`
}

type InputChannel struct {
	Name      string          `yaml:"name" json:"name"`
	Instance  int             `yaml:"instance" json:"instance"`
	Exchanges []InputExchange `yaml:"exchanges,mapstructure" json:"exchanges,mapstructure"`
}

/*
type InputConnections struct {
	Connections []InputConnection `yaml:"connections,mapstructure"`
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

type InputExchangeQueues struct {
	Exchanges []InputExcahnge `yaml:"exchanges,mapstructure" json:"exchanges,mapstructure"`
}

type InputQueueConsumers struct {
	QueueConsumers []InputQueueConsumer `yaml:"queueConsumers,mapstructure" json:"queueConsumers,mapstructure"`
}

type InputQueueConsumer struct {
	Name      string           `yaml:"name" json:"name"`
	Consumers []InputConsumers `yaml:"consumers,mapstructure" json:"consumers,mapstructure"`
}

type InputConsumers struct {
	ConsumerList []InputConsumer `yaml:"consumerList,mapstructure" json:"consumerList,mapstructure"`
}

type InputConsumer struct {
	Name string `yaml:"name" json:"name"`
}

type InputQueue struct {
	Name string `yaml:"name" json:"name"`
}

type InputExcahnge struct {
	Name   string       `yaml:"name" json:"name"`
	Queues []InputQueue `yaml:"queues,mapstructure" json:"queues,mapstructure"`
}

type InputExchangeTemplate struct {
	Name     string `yaml:"name" json:"name"`
	Template string `yaml:"template" json:"template"`
}

type InputExchange struct {
	Name string `yaml:"name" json:"name"`
}

/*
func getExchange(inputQueueConfig *InputQueueConfig, exchangeName string) InputExchangeTemplate {

	for _, e := range inputQueueConfig.Exchanges {
		if e.Name == exchangeName {
			for _, et := range inputQueueConfig.Templates.Exchanges {
				if et.Name == e.Template {
					return et
				}
			}
		}
	}

}
*/

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

func getRabbitMqArgs(args string) amqp.Table {
	table := make(amqp.Table)
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

func ValidateInputQueueConfig(inputQueueConfig *InputQueueConfig) {
	validateInputTemplates(inputQueueConfig.Templates)
	validateInputQueues(inputQueueConfig.Queues, inputQueueConfig.Templates.Queues)
	validateInputConsumers(inputQueueConfig.Consumers, inputQueueConfig.Templates.Consumers)
	validateInputQueueConsumers(inputQueueConfig.QueueConsumers, inputQueueConfig.Queues, inputQueueConfig.Consumers)
	validateInputExchange(inputQueueConfig.Exchanges, inputQueueConfig.Templates.Exchanges)
	validateInputExchangeQueues(inputQueueConfig.ExchangeQueues, inputQueueConfig.Exchanges, inputQueueConfig.Queues)
	validateInputConnection(inputQueueConfig.Connections, inputQueueConfig.Exchanges)

}

func validateInputTemplates(inputTemplate InputTemplate) {
	validateInputExchangeTemplate(inputTemplate.Exchanges)
	validateInputQueueTemplate(inputTemplate.Queues)
	validateInputConsumerTemplate(inputTemplate.Consumers)
}

func validateInputExchangeTemplate(exchanges []ExchangeTemplate) {
	mySet := make(utils.StringSet)
	if exchanges == nil || len(exchanges) <= 0 {
		log.Logger.Error("Invalid exchange configuration")
		panic("Invalid exchange configuration")
	}
	for _, e := range exchanges {
		if e.Name == "" {
			log.Logger.Error("Invalid exchange name")
			panic("Invalid exchange name")
		}
		utils.AddToSet(mySet, e.Name, "exchange template")
		if e.Type == "" {
			log.Logger.Error("Invalid exchange type")
			panic("Invalid exchange type")
		}
		err, found := utils.ValidateList(e.Type, "direct", "fanout", "topic", "headers", "default")
		if err != nil || found == false {
			log.Logger.Error("Invalid exchange type")
			panic("Invalid exchange type")
		}
		err, found = utils.ValidateList(e.Durable, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid exchange durable value")
			panic("Invalid exchange type durable value")
		}
		err, found = utils.ValidateList(e.AutoDelete, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid exchange autoDelete value")
			panic("Invalid exchange type autoDelete value")
		}
		err, found = utils.ValidateList(e.Internal, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid exchange internal value")
			panic("Invalid exchange type internal value")
		}
		err, found = utils.ValidateList(e.NoWait, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid exchange noWait value")
			panic("Invalid exchange noWait value")
		}
	}

}

func validateInputQueueTemplate(queues []QueueTemplate) {
	mySet := make(utils.StringSet)
	if queues == nil || len(queues) <= 0 {
		log.Logger.Error("Invalid queue configuration")
		panic("Invalid queue configuration")
	}
	for _, q := range queues {
		if q.Name == "" {
			log.Logger.Error("Invalid queue name")
			panic("Invalid queue name")
		}
		utils.AddToSet(mySet, q.Name, "queue template")
		err, found := utils.ValidateList(q.Durable, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid queue durable value")
			panic("Invalid queue type durable value")
		}
		err, found = utils.ValidateList(q.AutoDelete, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid queue autoDelete value")
			panic("Invalid queue type autoDelete value")
		}
		err, found = utils.ValidateList(q.Exclusive, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid queue exclusive value")
			panic("Invalid queue type exclusive value")
		}
		err, found = utils.ValidateList(q.NoWait, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid queue noWait value")
			panic("Invalid queue noWait value")
		}
	}

}

func validateInputConsumerTemplate(consumers []ConsumerTemplate) {
	mySet := make(utils.StringSet)
	if consumers == nil || len(consumers) <= 0 {
		log.Logger.Error("Invalid consumers configuration")
		panic("Invalid consumers configuration")
	}
	for _, q := range consumers {
		if q.Name == "" {
			log.Logger.Error("Invalid consumer name")
			panic("Invalid consumer name")
		}
		utils.AddToSet(mySet, q.Name, "consumer template")
		err, found := utils.ValidateList(q.AutoAck, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid consumer autoAck value")
			panic("Invalid consumer type autoAck value")
		}
		err, found = utils.ValidateList(q.Exclusive, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid consumer exclusive value")
			panic("Invalid consumer type exclusive value")
		}
		err, found = utils.ValidateList(q.NoLocal, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid consumer noLocal value")
			panic("Invalid consumer type noLocal value")
		}
		err, found = utils.ValidateList(q.NoWait, false, true)
		if err != nil || found == false {
			log.Logger.Error("Invalid consumer noWait internal value")
			panic("Invalid consumer type noWait value")
		}
	}

}

func validateInputQueues(Queues []InputQueueTemplate, queueTemplates []QueueTemplate) {
	found := false
	mySet := make(utils.StringSet)
	if Queues == nil || len(Queues) <= 0 {
		log.Logger.Error("Invalid consumers")
		panic("Invalid consumers")
	}
	for _, q := range Queues {
		if q.Name == "" {
			log.Logger.Error("Invalid queue name")
			panic("Invalid queue name")
		}
		utils.AddToSet(mySet, q.Name, "queues")
		if q.Template == "" {
			log.Logger.Error("Invalid queue template")
			panic("Invalid queue template")
		}
		for _, qt := range queueTemplates {
			if qt.Name == q.Template {
				found = true
				break
			}
		}
		if !found {
			log.Logger.Error("Queue template not found")
			panic("Queue template not found")
		}
	}
}

func validateInputConsumers(Consumers []InputConsumerTemplate, consumerTemplates []ConsumerTemplate) {
	found := false
	mySet := make(utils.StringSet)
	if Consumers == nil || len(Consumers) <= 0 {
		log.Logger.Error("Invalid consumers")
		panic("Invalid consumers")
	}
	for _, c := range Consumers {
		if c.Name == "" {
			log.Logger.Error("Invalid consumer name")
			panic("Invalid consumer name")
		}
		utils.AddToSet(mySet, c.Name, "consumers")
		if c.Template == "" {
			log.Logger.Error("Invalid consumer template")
			panic("Invalid consumer template")
		}
		for _, ct := range consumerTemplates {
			if ct.Name == c.Template {
				found = true
				break
			}
		}
		if !found {
			log.Logger.Error("Consumer template not found")
			panic("Consumer template not found")
		}
	}
}

func validateInputQueueConsumers(queueConsumers []InputQueueConsumers, queues []InputQueueTemplate, consumers []InputConsumerTemplate) {
	found := false
	queueSet := make(utils.StringSet)
	if queueConsumers == nil || len(queueConsumers) <= 0 {
		log.Logger.Error("Invalid queue consumers")
		panic("Invalid queue consumers")
	}
	for _, qcs := range queueConsumers {
		for _, qc := range qcs.QueueConsumers {
			if qc.Name == "" {
				log.Logger.Error("Invalid queue_consumer queue name")
				panic("Invalid queue consumer  queue name")
			}
			utils.AddToSet(queueSet, qc.Name, "queues in queue consumers")
			found = false
			for _, q := range queues {
				if qc.Name == q.Name {
					found = true
					break
				}
			}
			if !found {
				log.Logger.Error("Queue not found in Queue_Consumer")
				panic("Queue not found in Queue_Consumer")
			}
			if qc.Consumers == nil || len(qc.Consumers) <= 0 {
				log.Logger.Error("Invalid queue_consumer consumers")
				panic("Invalid queue_consumer consumers")
			}
			for _, q := range qc.Consumers {
				consumerSet := make(utils.StringSet)
				for _, ql := range q.ConsumerList {
					if ql.Name == "" {
						log.Logger.Error("Invalid queue_consumer consumer name")
						panic("Invalid queue_consumer consumer name")
					}
					utils.AddToSet(consumerSet, ql.Name, "consumers in queue consumers")
					found = false
					for _, c := range consumers {
						utils.AddToSet(consumerSet, c.Name, "consumers in queue consumers")
						if ql.Name == c.Name {
							found = true
							break
						}
					}
					if !found {
						log.Logger.Error("Consumer not found in Queue_Consumer")
						panic("Consumer not found in Queue_Consumer")
					}
				}
			}
		}
	}
}

func validateInputExchange(Exchanges []InputExchangeTemplate, exchangeTemplates []ExchangeTemplate) {
	found := false
	mySet := make(utils.StringSet)
	if Exchanges == nil || len(Exchanges) <= 0 {
		log.Logger.Error("Invalid exchange")
		panic("Invalid exchange")
	}
	for _, e := range Exchanges {
		if e.Name == "" {
			log.Logger.Error("Invalid exchange name")
			panic("Invalid exchange name")
		}
		utils.AddToSet(mySet, e.Name, "exchange")
		if e.Template == "" {
			log.Logger.Error("Invalid exchange template")
			panic("Invalid exchange template")
		}
		found = false
		for _, et := range exchangeTemplates {
			if et.Name == e.Template {
				found = true
				break
			}
		}
		if !found {
			log.Logger.Error("Exchange template not found")
			panic("Exchange template not found")
		}
	}
}

func validateInputExchangeQueues(exchangeQueues []InputExchangeQueues, exchanges []InputExchangeTemplate, queues []InputQueueTemplate) {
	found := false
	exchangeSet := make(utils.StringSet)
	if exchangeQueues == nil || len(exchangeQueues) <= 0 {
		log.Logger.Error("Invalid exchange queues")
		panic("Invalid exchange queues")
	}
	for _, eq := range exchangeQueues {
		for _, e := range eq.Exchanges {
			if e.Name == "" {
				log.Logger.Error("Invalid exchange queue name")
				panic("Invalid exchange queue name")
			}
			utils.AddToSet(exchangeSet, e.Name, "exchange")
			found = false
			for _, ex := range exchanges {
				if e.Name == ex.Name {
					found = true
					break
				}
			}
			if !found {
				log.Logger.Error("Exchange not found in exchange_queue")
				panic("Exchange not found in exchange_queue")
			}
			if e.Queues == nil || len(e.Queues) <= 0 {
				log.Logger.Error("Invalid exchange_queue queues")
				panic("Invalid exchange_queue queues")
			}
			for _, qs := range e.Queues {
				queueSet := make(utils.StringSet)
				utils.AddToSet(queueSet, qs.Name, "queue")
				found = false
				for _, q := range queues {
					if qs.Name == q.Name {
						found = true
						break
					}
				}
				if !found {
					log.Logger.Error("Queue not found in exchnage_queues")
					panic("Queue not found in exchnage_queues")
				}
			}
		}
	}
}

func validateInputConnection(connections []InputConnection, exchanges []InputExchangeTemplate) {
	found := false
	connectionSet := make(utils.StringSet)

	if connections == nil || len(connections) <= 0 {
		log.Logger.Error("Invalid connection configuration")
		panic("Invalid connection configuration")
	}
	for _, c := range connections {
		if c.Name == "" {
			log.Logger.Error("Invalid connection name")
			panic("Invalid connection name")
		}
		utils.AddToSet(connectionSet, c.Name, "connections")
		if c.ConnectionId == "" {
			log.Logger.Error("Invalid connection id")
			panic("Invalid connection id")
		}
		if c.ServerURL == "" {
			log.Logger.Error("Invalid connection server URL")
			panic("Invalid connection server URL")
		}
		if c.User == "" {
			log.Logger.Error("Invalid connection user")
			panic("Invalid connection user")
		}
		if c.Password == "" {
			log.Logger.Error("Invalid connection password")
			panic("Invalid connection password")
		}
		if c.Port == "" {
			log.Logger.Error("Invalid connection port")
			panic("Invalid connection port")
		}
		_, err := strconv.ParseFloat(c.Port, 64)
		if err != nil {
			log.Logger.Error("Invalid connection port")
			panic("Invalid connection port")
		}
		if c.UIPort == "" {
			log.Logger.Error("Invalid connection UI port")
			panic("Invalid connection UI port")
		}
		_, err = strconv.ParseFloat(c.UIPort, 64)
		if err != nil {
			log.Logger.Error("Invalid connection UI port")
			panic("Invalid connection port")
		}
		if c.Channels == nil || len(c.Channels) <= 0 {
			log.Logger.Error("Invalid connection channels")
			panic("Invalid connection channels")
		}
		channelSet := make(utils.StringSet)
		for _, cn := range c.Channels {
			if cn.Name == "" {
				log.Logger.Error("Invalid connection channel name")
				panic("Invalid connection channel name")
			}
			utils.AddToSet(channelSet, cn.Name, "channels")
			if cn.Instance > constants.MAX_CHANNELS {
				log.Logger.Error("Number of channels per connection above the limit")
				panic("Number of channels per connection above the limit")
			}
			if cn.Instance <= 0 {
				log.Logger.Error("At lease one channel needs to be defined")
				panic("At lease one channel needs to be defined")
			}
			if cn.Exchanges == nil || len(cn.Exchanges) <= 0 {
				log.Logger.Error("Invalid connection channel exchange")
				panic("Invalid connection channel exchange")
			}
			exchangeSet := make(utils.StringSet)
			for _, ex := range cn.Exchanges {
				if ex.Name == "" {
					log.Logger.Error("Invalid connection exchange name")
					panic("Invalid connection exchange name")
				}
				utils.AddToSet(exchangeSet, ex.Name, "exchanges")
				found = false
				for _, e := range exchanges {
					if ex.Name == e.Name {
						found = true
						break
					}
				}
				if !found {
					log.Logger.Error("Exchange not found in connection channel")
					panic("Exchange not found in connection channel")
				}
			}
		}
	}

}

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


func getQueueTemplates(inputQueueConfig *InputQueueConfig) map[string]*InputQueueTemplate {

	queues := make(map[string]*InputQueueTemplate)
	for _, t := range inputQueueConfig.Templates.Queues {
		inputQueueTemplate := InputQueueTemplate{
			Name:       t.Name,
			Durable:    t.Durable,
			AutoDelete: t.AutoDelete,
			Exclusive:  t.Exclusive,
			NoWait:     t.NoWait,
			Args:       t.Args,
		}
		queues[t.Name] = &inputQueueTemplate
	}

	return queues

}

func getConsumerTemplates(inputQueueConfig *InputQueueConfig) map[string]*InputConsumerTemplate {

	consumers := make(map[string]*InputConsumerTemplate)
	for _, t := range inputQueueConfig.Templates.Consumers {
		inputConsumerTemplate := InputConsumerTemplate{
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

//---------------------------    public   ---------------------------------------------------
func FillQueueTemplateValues(inputQueueConfig *InputQueueConfig) *InputQueueConfig {
	queueTemplates := getQueueTemplates(inputQueueConfig)
	consumerTemplates := getConsumerTemplates(inputQueueConfig)
	for _, con := range inputQueueConfig.Connections {
		for _, q := range con.Queues {
			if len(q.Template) > 0 {
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
