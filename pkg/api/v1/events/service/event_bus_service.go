package EventBusService

import (
	//"fmt"
	"encoding/json"
	"net/http"

	config "github.com/gilbarco-ai/event-bus/configuration"

	//"github.com/gilbarco-ai/event-bus/internal/common/constants"
	log "github.com/gilbarco-ai/event-bus/common/log"
	commonModel "github.com/gilbarco-ai/event-bus/common/model"
	sns "github.com/gilbarco-ai/event-bus/common/session"
	"github.com/gin-gonic/gin"
)

func InsertEvent(ctx *gin.Context) {
	var msgData config.MsgData
	if err := ctx.BindJSON(&msgData); err != nil {

		log.Logger.Error(err)
		//panic(err)
	}
	msg, err := json.Marshal(msgData)
	if err != nil {
		log.Logger.Error(err)
		//panic(err)
	}
	analyzeMsg(msgData)
	config.SendMsgByChannel(config.RabbitList[msgData.Station], string(msg))

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Insert event was finished succesfully"})
}

func UpdateEvent(ctx *gin.Context) {
	/*var newCronEvent EventBusModel.CronEventReq
	var session commonModel.Session

	session = sns.CreateSession(ctx)
	if err := ctx.BindJSON(&newCronEvent); err != nil {
		return
	}
	newCronEvent.CronId = ctx.Param("id")
	newCronEvent.UpdateDate = time.Now().UTC().Format(time.RFC3339)
	pgDao.UpdateCronEvent(session, newCronEvent)

	defer sns.CloseSession(session)

	UpdateCron(newCronEvent)
	*/
}

func GetEvents(ctx *gin.Context) {
	/*var newCronEvent EventBusModel.CronEventReq
	var session commonModel.Session
	session = sns.CreateSession(ctx)

	defer sns.CloseSession(session)

	CronEvents, err := pgDao.GetCronEvents(session, newCronEvent)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "CronEvent list not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, CronEvents)

	GetCrons()
	*/
}

func GetEventById(ctx *gin.Context) {
	/*	var session commonModel.Session

		id := ctx.Param("id")
		session = sns.CreateSession(ctx)

		defer sns.CloseSession(session)

		CronEvents, err := pgDao.GetCronEventById(session, id)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "CronEvent not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, CronEvents)

		GetCronById(id)
	*/
}

func DeleteEvent(ctx *gin.Context) {
	/*	var id string
		var session commonModel.Session

		session = sns.CreateSession(ctx)
		id = ctx.Param("id")
		pgDao.DeleteCronEvent(session, id)

		defer sns.CloseSession(session)

		DeleteCron(id)
	*/
}

func DeleteEventById(cronId string) {
	/*	var session commonModel.Session

		session = sns.CreateSession(nil)
		pgDao.DeleteCronEvent(session, cronId)

		defer sns.CloseSession(session)

		DeleteCron(cronId)
	*/
}

func Health(ctx *gin.Context) {

	var session commonModel.Session

	session = sns.CreateSession(ctx)

	defer sns.CloseSession(session)

	ctx.IndentedJSON(http.StatusOK, "Event bus service is healthy")

}

func analyzeMsg(msgData config.MsgData) {

	if len(msgData.MsgId) == 0 {
		log.Logger.Error("Invalid message id")
		//panic("Invalid message id")
	}
	if msgData.Station == "" {
		log.Logger.Error("Invalid station")
		//panic("Invalid station")
	}
	//config.AnalyzeMsg(msgData)

}
