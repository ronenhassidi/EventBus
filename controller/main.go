package main

// The main need to be identical for controller and agent ad copied to both directories /controller and /agent

import (
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

	// Load the queues
	generalViper := viper.GetViper(constants.GENERAL)
	appType := generalViper.GetString(constants.GENERAL_APP_TYPE) // controller \ agent

	appViper := viper.GetViper(appType)
	inputQueueConfig := &config.InputQueueConfig{}
	appViper.Unmarshal(inputQueueConfig)
	config.ValidateInputQueueConfig(inputQueueConfig)

	config.RabbitList = config.CreateQueueList(inputQueueConfig)
	//inputQueueConfig = config.FillQueueTemplateValues(inputQueueConfig)
	//	config.RabbitList = config.CreateQueueList(inputQueueConfig)

	//  ================================================== Need to be deleted =====================
	//	inputQueues := config.GetQueuesConfiguration(inputQueueConfig)
	//	testQueue(appType, inputQueues, inputQueueConfig)
	//  ================================================== Need to be deleted =====================

	// Release the queuee
	for _, r := range config.RabbitList {
		defer r.Connection.Close()
		defer r.Channel.Close()
	}

	// Initiate the HTTP server
	general := viper.GetViper("general")
	httpHost := general.GetString("general.httpHost")
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
