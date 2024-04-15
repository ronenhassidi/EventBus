package EventBusService

import (
	log "github.com/gilbarco-ai/event-bus/common/log"
	"github.com/gin-gonic/gin"
)

func StartServer(httpHost string) {

	const RoutePrefix = "api"

	router := gin.Default()
	err := BindRoutes(router.Group(RoutePrefix))
	if err != nil {
		log.Logger.Error(err)
	}

	log.Logger.Info(httpHost)
	router.Run(httpHost)
}
