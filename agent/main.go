package main

// The main need to be identical for controller and agent ad copied to both directories /controller and /agent

import (
	"os"

	viper "github.com/gilbarco-ai/event-bus/common/configuration/viper"
	"github.com/gilbarco-ai/event-bus/common/constants"
	log "github.com/gilbarco-ai/event-bus/common/log"
	config "github.com/gilbarco-ai/event-bus/configuration"
	pkg "github.com/gilbarco-ai/event-bus/pkg/api/common"
)

func main() {

	viper.LoadMainConfig()
	log.Init()
	log.Logger.Info("========================= Event bus started =========================")

	generalViper := viper.GetViper(constants.GENERAL)
	appType := generalViper.GetString(constants.GENERAL_APP_TYPE) // controller \ agent

	appViper := viper.GetViper(appType)
	inputQueueConfig := &config.InputQueueConfig{}
	appViper.Unmarshal(inputQueueConfig)
	config.ValidateInputQueueConfig(inputQueueConfig)

	config.CreateQueueList(inputQueueConfig)

	general := viper.GetViper("general")
	httpHost := general.GetString("general.httpHost")

	//  ================================================== Need to be deleted =====================
	// testQueue(appType, inputQueues, inputQueueConfig)
	//  ================================================== Need to be deleted =====================

	if appType == constants.CONTROLLER {

	} else if appType == constants.AGENT {
		queues := config.GetQueues(inputQueueConfig)

		//consumerList := config.GetConsumerList(inputQueueConfig)
		args := os.Args
		if len(args) == 2 {
			queueName := args[1]
			//consumerName := args[2]
			config.ConsumeMsg(config.RabbitList[queueName], queueName,
				queues[queueName].Consumers[0])
		} else {
			//config.RabbitList = config.CreateQueueList(inputQueueConfig)
			//for _, r := range rabbitList {
			//go consumer.ConsumeQueues(r)
			//}
		}
	}

	for _, r := range config.RabbitList {
		defer r.Connection.Close()
		defer r.Channel.Close()
	}

	pkg.StartServer(httpHost)

}

/*
func testQueue(appType string, inputQueues map[string]*config.QueueConfig, inputQueueConfig *config.InputQueueConfig) {

	if appType == constants.CONTROLLER {

		config.CreateQueues(inputQueues)
		msg := "444"
		config.SendMsgByChannel(config.RabbitList["station1"], msg)

		//msg = "5555"
		//config.SendMsgByChannel(rabbitList["station2"], msg)
	} else if appType == constants.AGENT {
		consumerList := config.GetConsumerList(inputQueueConfig)
		args := os.Args
		if len(args) == 2 {
			name := args[1]
			config.ConsumeMsg(config.RabbitList[name], consumerList[name], args[1])
		} else {
			rabbitList := config.CreateQueueList(inputQueueConfig)
			for _, r := range rabbitList {
				go consumer.ConsumeQueues(r)
			}
		}
	}

}

*/
