package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// download logic
// get group ,token -> search db -> get link -> expire ? yes -> get group , Link() -> get link  -> proxy
func Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		}
		matched, _ := regexp.MatchString("^/file/", c.Request.RequestURI)

		if !matched {
			return
		}

		fileId := strings.Split(c.Request.RequestURI, "/")[2]

		fileinfo, err := leveldb.GetOne[models.File](fileId)
		if err != nil {
			util.Response.ResponsFmt(c, http.StatusNotFound, 404, nil, "File Not Found")
			return
		}
		if fileinfo.Storage.Status != "work" {
			util.Response.Error(c, nil, "Storage Not Working")
			return
		}

		if fileinfo.Link.Header != nil {
			c.Request.Header = fileinfo.Link.Header
		}
		remote, err := url.Parse("http://" + fileinfo.Storage.DownloadAddr + "/" + fileinfo.Link.Path)

		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL = remote
		}
		proxy.ServeHTTP(c.Writer, c.Request)

		// c.Request.URL.Path = string(link.GetPath())

		// logger.Logger.Debug(c.Request)
		// // group, err := ldb.GetGroup(g)
		// HTTPProxy(c, "http", storage.ServerAddr)
	}
}

// HTTPProxy ,http 反向代理
func HTTPProxy(c *gin.Context, Scheme, Host string) bool {

	remote, err := url.Parse(Scheme + "://" + Host)
	if err != nil {
		return false
	}

	logger.Logger.Info("HTTPProxy: ", remote)
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(c.Writer, c.Request)
	return true
}
