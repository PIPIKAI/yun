package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/core/tracker/api"
)

func Router(g *gin.Engine) *gin.Engine {
	g.Use(api.Download())
	g.POST("/beforupload", api.BeforeUpload)
	g.POST("/upload", api.Upload)
	g.POST("/report-status", api.HanldeStorageServerReport)
	return g
}
