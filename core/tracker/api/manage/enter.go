// package manage ui apis
package manage

import (
	"time"

	"github.com/gin-gonic/gin"
)

// StartTime
var StartTime int64

// resourceRouter 静态资源配置
func resourceRouter(engine *gin.Engine) {

	html := NewHtmlHandler()
	engine = InitResource(engine)
	group := engine.Group("/ui")
	{
		group.GET("", html.Index)
	}
	engine.NoRoute(html.RedirectIndex)
}

// ManageRouter
//
//	@param g
func ManageRouter(g *gin.Engine) {
	StartTime = time.Now().Unix()
	resourceRouter(g)

	rg := g.Group("manage/")

	rg.GET("fileinfo", GetFileInfos)
	rg.GET("uploading", GetUploading)
	rg.GET("uploaded", GetUploaded)
	rg.GET("status", GetStatus)
	rg.DELETE("delsession", DelSession)
}
