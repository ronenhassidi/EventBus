package EventBusService

import (
	eventsV1 "github.com/gilbarco-ai/event-bus/pkg/api/v1/events"
	healthV1 "github.com/gilbarco-ai/event-bus/pkg/api/v1/health"
	eventsV2 "github.com/gilbarco-ai/event-bus/pkg/api/v2/events"
	healthV2 "github.com/gilbarco-ai/event-bus/pkg/api/v2/health"
	"github.com/gin-gonic/gin"
)

const VersionPrefixV1 = "v1"
const VersionPrefixV2 = "v2"

func BindRoutes(g *gin.RouterGroup) error {
	groupV1 := g.Group(VersionPrefixV1)
	healthV1.BindRoutes(groupV1)
	eventsV1.BindRoutes(groupV1)

	groupV2 := g.Group(VersionPrefixV2)
	healthV2.BindRoutes(groupV2)
	eventsV2.BindRoutes(groupV2)

	return nil
}
