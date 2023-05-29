package svc

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/core/tracker/api"
)

// Router
//
//	@param g
//	@return *gin.Engine
func Router(g *gin.Engine) *gin.Engine {
	pprof.Register(g)
	g.Use(api.Download())
	rg := g.Group("/api")
	{
		rg.POST("/beforupload", api.BeforeUpload)
		rg.POST("/reupload", api.ReUploadSession)
		rg.POST("/upload", api.Upload)
		rg.POST("/create", api.Create)
	}
	g.POST("/report-status", api.HanldeStorageServerReport)
	g.POST("/report-sync", api.HandleSync)

	return g
}
