package EventBusService

import (
	eventService "github.com/gilbarco-ai/event-bus/pkg/api/v1/events/service"
	"github.com/gin-gonic/gin"
)

func BindRoutes(g *gin.RouterGroup) {
	g.POST("/event-bus", eventService.InsertEvent)
	g.GET("/events-bus/", eventService.GetEvents)
	g.GET("/event-bus/:id", eventService.GetEventById)
	g.PATCH("/event-bus/:id", eventService.UpdateEvent)
	g.DELETE("/event-bus/:id", eventService.DeleteEvent)

}
