package configuration

import (
	"strconv"

	"github.com/gilbarco-ai/event-bus/common/constants"
	log "github.com/gilbarco-ai/event-bus/common/log"
	utils "github.com/gilbarco-ai/event-bus/common/utils"
)

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
		if err != nil || !found {
			log.Logger.Error("Invalid exchange type")
			panic("Invalid exchange type")
		}
		err, found = utils.ValidateList(e.Durable, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid exchange durable value")
			panic("Invalid exchange type durable value")
		}
		err, found = utils.ValidateList(e.AutoDelete, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid exchange autoDelete value")
			panic("Invalid exchange type autoDelete value")
		}
		err, found = utils.ValidateList(e.Internal, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid exchange internal value")
			panic("Invalid exchange type internal value")
		}
		err, found = utils.ValidateList(e.NoWait, false, true)
		if err != nil || !found {
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
		if err != nil || !found {
			log.Logger.Error("Invalid queue durable value")
			panic("Invalid queue type durable value")
		}
		err, found = utils.ValidateList(q.AutoDelete, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid queue autoDelete value")
			panic("Invalid queue type autoDelete value")
		}
		err, found = utils.ValidateList(q.Exclusive, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid queue exclusive value")
			panic("Invalid queue type exclusive value")
		}
		err, found = utils.ValidateList(q.NoWait, false, true)
		if err != nil || !found {
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
		if err != nil || !found {
			log.Logger.Error("Invalid consumer autoAck value")
			panic("Invalid consumer type autoAck value")
		}
		err, found = utils.ValidateList(q.Exclusive, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid consumer exclusive value")
			panic("Invalid consumer type exclusive value")
		}
		err, found = utils.ValidateList(q.NoLocal, false, true)
		if err != nil || !found {
			log.Logger.Error("Invalid consumer noLocal value")
			panic("Invalid consumer type noLocal value")
		}
		err, found = utils.ValidateList(q.NoWait, false, true)
		if err != nil || !found {
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
	for _, eq := range queueConsumers {
		if eq.Name == "" {
			log.Logger.Error("Invalid queue name")
			panic("Invalid queue name")
		}
		found = false
		for _, ex := range queues {
			if eq.Name == ex.Name {
				found = true
				break
			}
		}
		if !found {
			log.Logger.Error("Queue not found in queue_consumer")
			panic("Queue not found in queue_consumer")
		}
		utils.AddToSet(queueSet, eq.Name, "queue")
		consumerSet := make(utils.StringSet)
		for _, q := range eq.Consumers {
			if q.Name == "" {
				log.Logger.Error("Invalid queue consumer name")
				panic("Invalid queue consumer name")
			}
			found = false
			for _, qs := range consumers {
				if qs.Name == q.Name {
					found = true
					break
				}
			}
			if !found {
				log.Logger.Error("Consumer not found in exchnage_consumers")
				panic("Consumer not found in exchnage_consumers")
			}
			utils.AddToSet(consumerSet, q.Name, "consumer")
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
		if eq.Name == "" {
			log.Logger.Error("Invalid exchange name")
			panic("Invalid exchange name")
		}
		found = false
		for _, ex := range exchanges {
			if eq.Name == ex.Name {
				found = true
				break
			}
		}
		if !found {
			log.Logger.Error("Exchange not found in exchange_queue")
			panic("Exchange not found in exchange_queue")
		}
		utils.AddToSet(exchangeSet, eq.Name, "exchange")
		queueSet := make(utils.StringSet)
		for _, q := range eq.Queues {
			if q.Name == "" {
				log.Logger.Error("Invalid exchange queue name")
				panic("Invalid exchange queue name")
			}
			found = false
			for _, qs := range queues {
				if qs.Name == q.Name {
					found = true
					break
				}
			}
			if !found {
				log.Logger.Error("Queue not found in exchnage_queues")
				panic("Queue not found in exchnage_queues")
			}
			utils.AddToSet(queueSet, q.Name, "queue")
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
