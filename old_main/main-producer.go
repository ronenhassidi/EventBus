package main

/*
import (
	"os"

//	config "github.com/gilbarco-ai/event-bus/configuration"
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
	args := os.Args
	if len(args) != 2 {
		panic("Queue is mandatory argument")
	}
	if inputQueues[args[1]] == nil {
		panic("Queue does not exist")
	}
	config.SendMsg(inputQueues[args[1]], "Ronen")

}
*/
