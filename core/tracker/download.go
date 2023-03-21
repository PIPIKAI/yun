package tracker

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
)

func (t *tracker) Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		}
		matched, _ := regexp.MatchString("^/group/", c.Request.RequestURI)

		if !matched {
			return
		}

		g := strings.Split(c.Request.RequestURI, "/")[2]

		token := strings.Split(c.Request.RequestURI, "/")[3]
		if g == "" {
			util.Response.ResponsFmt(c, http.StatusNotFound, 404, nil, "")
			return
		}

		ldb, err := leveldb.NewLDB(consts.Group_Storage_DB)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}

		logger.Logger.Info("g ", g)
		logger.Logger.Info("token ", token)

		group, err := ldb.GetGroup(g)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
		storage, err := t.SelectStorage(c, *group)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}

		fdb, err := leveldb.NewLDB(consts.File_List_DB)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
		path, err := fdb.Do(token)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
		c.Request.URL.Path = string(path)

		// group, err := ldb.GetGroup(g)
		// c.Redirect(http.StatusPermanentRedirect, "http://"+storage.ServerAddr+string(path))
		t.HTTPProxy(c, "http", storage.ServerAddr)
	}
}

// HTTPProxy ,http 反向代理
func (t *tracker) HTTPProxy(c *gin.Context, Scheme, Host string) bool {
	remote, err := url.Parse(Scheme + "://" + Host)
	if err != nil {
		return false
	}

	logger.Logger.Info("HTTPProxy: ", remote)
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(c.Writer, c.Request)
	return true
}
