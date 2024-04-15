package session

import (
	"log"

	config "github.com/gilbarco-ai/event-bus/common/configuration/viper"
	dbCore "github.com/gilbarco-ai/event-bus/common/db_core/service"
	commonModel "github.com/gilbarco-ai/event-bus/common/model"
	"github.com/gin-gonic/gin"
)

func CreateSession(ctx *gin.Context) commonModel.Session {
	var con dbCore.Connection
	var session commonModel.Session

	if ctx != nil {
		session.GinCtx = ctx
	}

	dbVP := config.GetViper("db")
	err := dbVP.Unmarshal(&con)
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbCore.OpenDB(con)
	if err == nil {
		session.DB = db
		ctx, tx, err := dbCore.CreateTransaction(session.DB)
		if err != nil {
			log.Fatal(err)
		}
		session.Ctx = ctx
		session.Tx = tx

	} else {
		log.Fatal(err)
	}
	return session
}

func CloseSession(session commonModel.Session) {
	if session.DB != nil {
		dbCore.CloseDB(session.DB)
	}
}
