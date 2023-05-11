package util

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// RestyClient
var RestyClient = NewRestyClient()

var noRedirectClient *resty.Client
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
var defaultTimeout = time.Second * 30

func init() {
	noRedirectClient = resty.New().SetRedirectPolicy(
		resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}),
	)
	noRedirectClient.SetHeader("user-agent", userAgent)
}

func NewRestyClient() *resty.Client {
	client := resty.New().
		SetHeader("user-agent", userAgent).
		SetRetryCount(3).
		SetTimeout(defaultTimeout)
	return client
}
