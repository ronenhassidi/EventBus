package Health

import (
	"fmt"
	"net/http"

	log "github.com/gilbarco-ai/event-bus/common/log"
	config "github.com/gilbarco-ai/event-bus/configuration"
	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {

	k := []config.KeyValue{
		{
			Key:   "a",
			Value: "b",
		},
		{
			Key:   "c",
			Value: "d",
		},
	}
	res := config.ApiResponse{
		Status:        "Success",
		Message:       "Event bus v1 service is healthy",
		KeyValueArray: k,
	}

	//ctx.IndentedJSON(http.StatusOK, "Event bus v1 service is healthy")
	ctx.IndentedJSON(http.StatusOK, res)
	log.Logger.Info("Event bus v1 service is healthy")
	fmt.Println("Event bus v1 service is healthy")
}
