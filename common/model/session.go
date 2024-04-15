package commonModel

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Session struct {
	DB     *sql.DB         `json"db"`
	Ctx    context.Context `json"ctx"`
	GinCtx *gin.Context    `json"ginCtx"`
	Tx     *sql.Tx         `json"tx"`
}
