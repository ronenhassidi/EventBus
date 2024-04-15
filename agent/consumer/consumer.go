package consumer

import (
	"encoding/json"
	"fmt"

	commands "github.com/gilbarco-ai/event-bus/commands/http"
	"github.com/gilbarco-ai/event-bus/common/constants"
	log "github.com/gilbarco-ai/event-bus/common/log"

	config "github.com/gilbarco-ai/event-bus/configuration"
)

func ConsumeQueues(rabbit config.Rabbit) {

	msgs, err := rabbit.Channel.Consume(
		rabbit.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
			msgData := config.MsgData{}
			err := json.Unmarshal([]byte(d.Body), &msgData)
			if err != nil {
				fmt.Println("Error:", err)
				log.Logger.Error(err)
			}
			executeMsg(msgData)
			fmt.Println(msgData)
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println(" [*] - Waitint for messages")
	<-forever

}

func executeMsg(msgData config.MsgData) {

	switch msgType := msgData.MsgType; msgType {
	case constants.API_TYPE:
		executeApiMsg(msgData)
	case constants.DB_TYPE:
		executeDbMsg(msgData)
	case constants.FS_TYPE:
		executeFsMsg(msgData)
	default:
		err := "Invalid message type"
		log.Logger.Error(err)
	}
}

func executeApiMsg(msgData config.MsgData) {
	fmt.Println("execute API")
	commands.HttpClient(msgData.ApiMsg)
}

func executeDbMsg(msgData config.MsgData) {
	fmt.Println("execute DB command")
}

func executeFsMsg(msgData config.MsgData) {
	fmt.Println("execute FS command")
}
