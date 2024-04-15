package ppxlog

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	log "github.com/gilbarco-ai/event-bus/common/log"
)

type PPXLog struct {
	Timestamp       string `yaml:"timestamp" json:"timestamp"`
	Loglevel        string `yaml:"log.level" json:"log.level"`
	Message         string `yaml:"message" json:"message"`
	MessageTemplete string `yaml:"message.templete" json:"message.templete"`
	LoggerName      string `yaml:"logger.name" json:"logger.name"`
	ThreadId        string `yaml:"thread.id" json:"thread.id"`
	ProccessId      string `yaml:"proccess.id" json:"proccess.id"`
}

func Error(msg string) {
	// Log some messages with custom fields
	PPXLog := PPXLog{
		Timestamp:       strconv.FormatInt(getUnixPicks(), 10),
		Loglevel:        "ERROR",
		Message:         msg,
		MessageTemplete: "Code Certificate",
		LoggerName:      "Code Certificate Logger",
		ThreadId:        "0",
		ProccessId:      "0",
	}
	jsonBytes, err := json.Marshal(PPXLog)
	if err != nil {
		log.Logger.Error("Error:", err)
		return
	}
	jsonString := string(jsonBytes)
	jsonString = strings.ReplaceAll(jsonString, `\`, "")

	log.Logger.Error(string(jsonString))
}

func Info(msg string) {
	// Log some messages with custom fields
	PPXLog := PPXLog{
		Timestamp:       strconv.FormatInt(getUnixPicks(), 10),
		Loglevel:        "INFO",
		Message:         msg,
		MessageTemplete: "Code Certificate",
		LoggerName:      "Code Certificate Logger",
		ThreadId:        "0",
		ProccessId:      "0",
	}
	jsonBytes, err := json.Marshal(PPXLog)
	if err != nil {
		log.Logger.Info("Info:", err)
		return
	}
	jsonString := string(jsonBytes)
	jsonString = strings.ReplaceAll(jsonString, `\`, "")

	log.Logger.Error(string(jsonString))
}

func getUnixPicks() int64 {
	currentTime := time.Now()
	unixTimestampSeconds := currentTime.Unix()
	return unixTimestampSeconds
}
