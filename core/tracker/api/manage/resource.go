package manage

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/resources"
)

func InitResource(engine *gin.Engine) *gin.Engine {
	fsys, err := fs.Sub(resources.Static, "html/static")
	if err != nil {
		panic(err)
	}
	engine.StaticFS("/static", http.FS(fsys))
	return engine
}

type HtmlHandler struct{}

func NewHtmlHandler() *HtmlHandler {
	return &HtmlHandler{}
}

// RedirectIndex 重定向
func (h *HtmlHandler) RedirectIndex(c *gin.Context) {
	c.Redirect(http.StatusFound, "/ui")
}

func (h *HtmlHandler) Index(c *gin.Context) {
	c.Header("content-type", "text/html;charset=utf-8")
	c.String(200, string(resources.Html))
}
