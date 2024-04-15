package health

import (
	healthService "github.com/gilbarco-ai/event-bus/pkg/api/v2/health/service"
	"github.com/gin-gonic/gin"
)

func BindRoutes(g *gin.RouterGroup) {
	g.GET("/health", healthService.HealthCheck)
}
