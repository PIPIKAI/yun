package models

import (
	"io"
	"net/http"
	"time"
)

type MLink struct {
	Scheme    string
	Path      string
	Header    http.Header `json:"header"` // needed header
	ExpriedAt time.Time
}

func (l MLink) GetScheme() string {
	return l.Scheme
}
func (l MLink) GetPath() string {
	return l.Path
}
func (l MLink) GetHeader() http.Header {
	return l.Header
}

var LinkDB = "Link"

// k : filemetaID v : Link
type Link struct {
	Scheme     string         `json:"scheme"`
	Path       string         `json:"path"`
	Header     http.Header    `json:"header"` // needed header
	Data       io.ReadCloser  // return file reader directly
	Expiration *time.Duration // url expiration time
}
