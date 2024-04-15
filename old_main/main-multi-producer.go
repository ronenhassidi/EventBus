package main

/*
import (
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
	inputQueueConfig := &config.InputQueueConfig{}
	queueViper.Unmarshal(inputQueueConfig)

	//inputQueues := config.GetQueuesConfiguration(inputQueueConfig)
	//config.CreateQueues(inputQueues)

	rabbitList := config.CreateQueueList(inputQueueConfig)
	//len := len(rabbitList)
	for i := 0; i < 1000; i++ {
		//randomInt := rand.Intn(len)
		//msg := strconv.Itoa(i + 1)
		//config.SendMsgByChannel(rabbitList[randomInt], msg)
		sendToQueues(rabbitList, i)
	}

	for _, r := range rabbitList {
		defer r.Connection.Close()
		defer r.Channel.Close()
	}
}

func sendToQueues(rabbitList []config.Rabbit, msgNum int) {
	for i := 0; i < len(rabbitList); i++ {
		msg := "Queue No-" + strconv.Itoa(i+1) + " Msg-" + strconv.Itoa(msgNum)
		config.SendMsgByChannel(rabbitList[i], msg)
	}
}
*/