package main

/*
import (
	"fmt"
	"math/rand"
	"strconv"

	config "github.com/gilbarco-ai/event-bus/configuration"
	viper "github.com/gilbarco-ai/event-bus/internal/common/configuration/viper"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	queueName  string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func main() {

	//  viper.LoadMainConfig()
	//  queue := viper.GetViper("queue")
	//  queueConfig := &config.QueueConfig{}
	//  queue.Unmarshal(queueConfig)
	//	fmt.Println(queueConfig.Connections[0].ConnectionId)
	//	fmt.Println(queueConfig.Connections[0].Queues[0].Name)
	//	fmt.Println(queueConfig.Connections[0].Queues[0].Consumers[0].Name)

	viper.LoadMainConfig()
	queueViper := viper.GetViper("queue")
	inputQueueConfig := &config.InputQueueConfig{}
	queueViper.Unmarshal(inputQueueConfig)



	queues := config.GetQueuesConfiguration(inputQueueConfig)
	config.CreateQueues(queues)

	var rabbit [4]Rabbit
	for i := 0; i < 4; i++ {
		rabbit[i].queueName = "queue" + strconv.Itoa(i+1)

		if queues[rabbit[i].queueName] == nil {
			fmt.Println("queue not found")
			return
		}


		conn, err := queues[rabbit[i].queueName].Connections.GetConnection()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		rabbit[i].Connection = conn
		defer rabbit[i].Connection.Close()
		fmt.Println("Succesfully connected to RabbitMQ")

		ch, err := rabbit[i].Connection.Channel()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		rabbit[i].Channel = ch
		defer rabbit[i].Channel.Close()
	}
	for i := 0; i < 100000; i++ {
		randomInt := rand.Intn(4)
		msg := strconv.Itoa(i + 1)
		config.SendMsgByChannel(rabbit[randomInt].Channel, rabbit[randomInt].queueName, msg)
	}

}




*/
