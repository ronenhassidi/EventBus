package Health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {

	ctx.IndentedJSON(http.StatusOK, "Event bus v2 service is healthy")

}
