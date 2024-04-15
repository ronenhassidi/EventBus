package main

/*
import (
	"os"

	config "github.com/gilbarco-ai/event-bus/configuration"
	viper "github.com/gilbarco-ai/event-bus/internal/common/configuration/viper"
)

func main() {
	viper.LoadMainConfig()
	queueViper := viper.GetViper("queue")
	inputQueueConfig := &config.InputQueueConfig{}
	queueViper.Unmarshal(inputQueueConfig)

	queues := config.GetQueuesConfiguration(inputQueueConfig)
	args := os.Args
	if len(args) != 2 {
		panic("Queue is mandatory argument")
	}
	if queues[args[1]] == nil {
		panic("Queue does not exist")
	}
	config.ConsumeMsg(queues[args[1]], args[1])

}
*/
