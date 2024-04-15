package main

/*
import (
	"math/rand"
	"strconv"

	config "github.com/gilbarco-ai/event-bus/configuration"
	viper "github.com/gilbarco-ai/event-bus/internal/common/configuration/viper"
)

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
	queueInputConfig := &config.QueueInputConfig{}
	queueViper.Unmarshal(queueInputConfig)

	inputQueues := config.GetQueuesConfiguration(queueInputConfig)
	config.CreateQueues(inputQueues)

	rabbit := make([]config.Rabbit, 4)
	for i := 0; i < 4; i++ {
		queueName := "queue" + strconv.Itoa(i+1)
		r := config.GetRabbit(inputQueues[queueName], queueName)
		rabbit[i] = r
	}
	for i := 0; i < 1000000; i++ {
		randomInt := rand.Intn(4)
		msg := strconv.Itoa(i + 1)
		config.SendMsgByChannel(rabbit[randomInt], msg)
	}
	for _, r := range rabbit {
		defer r.Connection.Close()
		defer r.Channel.Close()
	}
}
*/
