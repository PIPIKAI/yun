// package
package api

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/logger"
)

// Proxy
//
//	@return gin.HandlerFunc
func Proxy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matched, _ := regexp.MatchString("^/http", ctx.Request.RequestURI)
		if !matched {
			return
		}

		req_url := ctx.Request.RequestURI[1:]
		logger.Logger.Info(req_url)

		ctx.Redirect(http.StatusMovedPermanently, req_url)
	}
}
