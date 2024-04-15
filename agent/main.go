package main

/*
// The main need to be identical for controller and agent ad copied to both directories /controller and /agent

import (
	"os"

	consumer "github.com/gilbarco-ai/event-bus/agent/consumer"
	config "github.com/gilbarco-ai/event-bus/configuration"
	viper "github.com/gilbarco-ai/event-bus/internal/common/configuration/viper"
	"github.com/gilbarco-ai/event-bus/internal/common/constants"
	logger "github.com/gilbarco-ai/event-bus/internal/common/log"
	pkg "github.com/gilbarco-ai/event-bus/pkg/api/common"
)

func main() {

	logger.Init("../logs/app.log", false)
	logger.Info.Println("========================= Event bus started =========================")
	viper.LoadMainConfig()

	generalViper := viper.GetViper(constants.GENERAL)
	appType := generalViper.GetString(constants.GENERAL_APP_TYPE) // controller \ agent

	appViper := viper.GetViper(appType)
	inputQueueConfig := &config.InputQueueConfig{}
	appViper.Unmarshal(inputQueueConfig)

	inputQueueConfig = config.FillQueueTemplateValues(inputQueueConfig)
	config.RabbitList = config.CreateQueueList(inputQueueConfig)
	//inputQueues := config.GetQueuesConfiguration(inputQueueConfig)

	general := viper.GetViper("general")
	httpHost := general.GetString("general.httpHost")

	//  ================================================== Need to be deleted =====================
	// testQueue(appType, inputQueues, inputQueueConfig)
	//  ================================================== Need to be deleted =====================

	if appType == constants.CONTROLLER {

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

	for _, r := range config.RabbitList {
		defer r.Connection.Close()
		defer r.Channel.Close()
	}

	pkg.StartServer(httpHost)

}

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
