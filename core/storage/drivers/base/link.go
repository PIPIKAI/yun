package base

import "net/http"

type Link struct {
	Scheme string
	Path   string
	Header http.Header `json:"header"` // needed header
}

func (l Link) GetScheme() string {
	return l.Scheme
}
func (l Link) GetPath() string {
	return l.Path
}
func (l Link) GetHeader() http.Header {
	return l.Header
}
